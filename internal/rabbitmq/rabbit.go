package rabbitmq

import (
	"TJ-project/models"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func Connect(rabbitMQURL string) error {
	var err error
	conn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		return err
	}

	ch, err = conn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		"fleet.events",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("Connected to RabbitMQ successfully.")
	return nil
}

func Close() {
	if err := ch.Close(); err != nil {
		log.Printf("Error closing RabbitMQ channel: %v", err)
	}
	if err := conn.Close(); err != nil {
		log.Printf("Error closing RabbitMQ connection: %v", err)
	}
}

func PublishGeofenceEvent(loc models.Location) {
	body, err := json.Marshal(map[string]interface{}{
		"vehicle_id": loc.VehicleID,
		"event":      "geofence_entry",
		"location": map[string]float64{
			"latitude":  loc.Latitude,
			"longitude": loc.Longitude,
		},
		"timestamp": loc.Timestamp,
	})
	if err != nil {
		log.Printf("Error marshaling geofence event: %v", err)
		return
	}

	err = ch.Publish(
		"fleet.events",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Error publishing geofence event: %v", err)
	}
}
