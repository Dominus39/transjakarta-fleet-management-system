package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Location struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://mosquitto:1884")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		vehicleID := generateRandomVehicleID()
		loc := Location{
			VehicleID: vehicleID,
			Latitude:  -6.2088 + rand.Float64()*0.01, // Random near Jakarta
			Longitude: 106.8456 + rand.Float64()*0.01,
			Timestamp: time.Now().Unix(),
		}

		payload, _ := json.Marshal(loc)
		topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)
		client.Publish(topic, 0, false, payload)

		fmt.Println("Published:", string(payload))
		time.Sleep(2 * time.Second)
	}
}

func generateRandomVehicleID() string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(900) + 100 // Generate a random number between 100 and 999
	return fmt.Sprintf("vehicle_%d", randomNum)
}
