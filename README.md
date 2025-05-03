# Transjakarta Fleet Management System

This is a fleet management system for Transjakarta, built using **Go (Golang)**, **MQTT**, **PostgreSQL**, **RabbitMQ**, and **Docker**. The application provides APIs to manage vehicle location tracking, history, geofence events, and more.

## Features

- Track real-time vehicle locations.
- Retrieve vehicle location history for a given time range.
- Publish geofence events when a vehicle enters a predefined geofence area.
- Containerized using Docker Compose for easy deployment.

## Requirements

Before running this application, you need to have the following installed:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go (Golang)](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [RabbitMQ](https://www.rabbitmq.com/download.html)
- [Postman](https://www.postman.com/) (for API testing)

## Getting Started

1. **Clone the repository:**

   Clone this repository to your local machine using:
   
   git clone https://github.com/Dominus39/transjakarta-fleet-management-system.git  
   
2. **Set up the environment:**

   Create a .env file in the root of the project and define the necessary environment variables (like RabbitMQ URL, MQTT broker, etc.).

   Example .env:

   DATABASE_URL=postgres://postgres:postgres@postgres:5432/fleet?sslmode=disable
   MQTT_BROKER=tcp://mosquitto:1883
   MQTT_TOPIC=/fleet/vehicle/+/location
   RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
   RABBITMQ_EXCHANGE=fleet.events
   RABBITMQ_QUEUE=geofence_alerts
   GEOFENCE_LAT=-6.2088
   GEOFENCE_LON=106.8456
   GEOFENCE_RADIUS=50
 
3. **Docker Setup:**

    Run the following command to build and start the services using Docker Compose:

    docker-compose up --build
    This will start the PostgreSQL database, RabbitMQ, and your Go application inside Docker containers.

4. **Start the application locally:**

   If you prefer not to use Docker, you can run the application locally with:

   go run main.go
   This will start the server on the port defined in your .env file (default is 8080).

## API Endpoints
1. **Get Latest Vehicle Location**
   GET /vehicles/:vehicle_id/location
    
   Fetches the latest location of the specified vehicle.
    
   Example Request:

   GET http://localhost:8080/vehicles/vehicle_363/location
   
   Response:
   {
     "vehicle_id": "vehicle_363",
     "latitude": -6.202012058733707,
     "longitude": 106.84643041830446,
     "timestamp": 1746250904
   }
   
2. **Get Vehicle Location History**
   GET /vehicles/:vehicle_id/history

   Retrieves the location history for the specified vehicle between a start and end timestamp.

   Example Request:

   GET http://localhost:8080/vehicles/vehicle_363/history?start=1746250000&end=1746250904

   Response:
   [
     {
       "vehicle_id": "vehicle_363",
       "latitude": -6.202012058733707,
       "longitude": 106.84643041830446,
       "timestamp": 1746250904
     }
   ]
   
3. **Geofence Event (When a vehicle enters the geofence)**
   Event: A message is sent to RabbitMQ when a vehicle enters the geofence area.

   Example: If a vehicle enters the geofence area (defined by the latitude, longitude, and radius), a RabbitMQ message is published.
