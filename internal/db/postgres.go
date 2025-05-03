package db

import (
	"TJ-project/config"
	"TJ-project/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := config.AppConfig.DBURL
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL successfully.")
	return nil
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

func SaveLocation(loc models.Location) error {
	_, err := DB.Exec("INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp) VALUES ($1, $2, $3, $4)",
		loc.VehicleID, loc.Latitude, loc.Longitude, loc.Timestamp)
	return err
}

func GetLatestLocation(vehicleID string) (models.Location, error) {
	var loc models.Location

	query := `
		SELECT vehicle_id, latitude, longitude, timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`

	row := DB.QueryRow(query, vehicleID)

	err := row.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
	if err != nil {
		return loc, err
	}

	return loc, nil
}

func GetLocationHistory(vehicleID string, start, end int64) ([]models.Location, error) {
	query := `
		SELECT vehicle_id, latitude, longitude, timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1 AND timestamp BETWEEN $2 AND $3
		ORDER BY timestamp ASC
	`

	log.Printf("Query: %s, VehicleID: %s, Start: %d, End: %d", query, vehicleID, start, end)

	rows, err := DB.Query(query, vehicleID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.Location
	for rows.Next() {
		var loc models.Location

		if err := rows.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp); err != nil {
			return nil, err
		}

		history = append(history, loc)
	}

	return history, nil
}
