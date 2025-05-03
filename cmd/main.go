package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"TJ-project/config"
	"TJ-project/internal/api"
	"TJ-project/internal/db"
	"TJ-project/internal/mqtt"
	"TJ-project/internal/rabbitmq"
)

func main() {
	// 1. Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// 2. Load app configuration
	config.Load()

	// 3. Connect to PostgreSQL
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// 4. Connect to RabbitMQ
	if err := rabbitmq.Connect(config.AppConfig.RabbitMQURL); err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitmq.Close()

	// 5. Start MQTT
	go mqtt.StartMQTTBroker(config.AppConfig.MQTTBroker, config.AppConfig.MQTTTopic)

	// 6. Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	api.StartServer(":" + port)
}
