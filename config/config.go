package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBURL          string
	MQTTBroker     string
	MQTTTopic      string
	RabbitMQURL    string
	RabbitMQEx     string
	RabbitMQQueue  string
	GeofenceLat    float64
	GeofenceLon    float64
	GeofenceRadius float64
}

var AppConfig Config

func waitForDB() {
	for i := 0; i < 5; i++ {
		err := checkDBConnection()
		if err == nil {
			return
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Failed to connect to DB after multiple attempts")
	os.Exit(1)
}

func checkDBConnection() error {
	db, err := sql.Open("postgres", AppConfig.DBURL)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func Load() {
	lat, _ := strconv.ParseFloat(getEnv("GEOFENCE_LAT", "-6.2088"), 64)
	lon, _ := strconv.ParseFloat(getEnv("GEOFENCE_LON", "106.8456"), 64)
	radius, _ := strconv.ParseFloat(getEnv("GEOFENCE_RADIUS", "50"), 64)

	AppConfig = Config{
		DBURL:          getEnv("DATABASE_URL", "postgres://postgres:postgres@postgres:5432/fleet?sslmode=disable"),
		MQTTBroker:     getEnv("MQTT_BROKER", "tcp://mosquitto:1884"),
		MQTTTopic:      getEnv("MQTT_TOPIC", "/fleet/vehicle/+/location"),
		RabbitMQURL:    getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/"),
		RabbitMQEx:     getEnv("RABBITMQ_EXCHANGE", "fleet.events"),
		RabbitMQQueue:  getEnv("RABBITMQ_QUEUE", "geofence_alerts"),
		GeofenceLat:    lat,
		GeofenceLon:    lon,
		GeofenceRadius: radius,
	}
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
