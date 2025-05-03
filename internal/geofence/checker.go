package geofence

import (
	"TJ-project/internal/rabbitmq"
	"TJ-project/models"
	"math"
)

var centerLat = -6.2000
var centerLon = 106.8450
var radius = 50.0 // meters

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func CheckGeofence(loc models.Location) {
	if haversine(centerLat, centerLon, loc.Latitude, loc.Longitude) <= radius {
		rabbitmq.PublishGeofenceEvent(loc)
	}
}
