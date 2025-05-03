package mqtt

import (
	"TJ-project/internal/db"
	"TJ-project/internal/geofence"
	"TJ-project/models"
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMQTTBroker(broker, topic string) {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)

	maxRetries := 5
	var connectErr error

	for i := 1; i <= maxRetries; i++ {
		token := client.Connect()
		token.Wait()
		connectErr = token.Error()

		if connectErr == nil {
			log.Println("Connected to MQTT broker:", broker)
			break
		}

		log.Printf("Failed to connect to MQTT broker (%s): %v. Retry %d/%d in 5s...", broker, connectErr, i, maxRetries)
		time.Sleep(5 * time.Second)
	}

	if connectErr != nil {
		log.Fatalf("Could not connect to MQTT broker after %d attempts: %v", maxRetries, connectErr)
	}

	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		var loc models.Location
		if err := json.Unmarshal(msg.Payload(), &loc); err != nil {
			log.Println("Invalid payload:", err)
			return
		}

		loc.Timestamp = time.Now().Unix()

		if err := db.SaveLocation(loc); err != nil {
			log.Println("Error saving location:", err)
		}

		go geofence.CheckGeofence(loc)
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", topic, token.Error())
	}
}
