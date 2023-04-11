package main

import (
	"log"

	"github.com/SinisterSup/messageQueuing/internal/app/api"
	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/rabbitmq"
	"github.com/gin-gonic/gin"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.LoadConfig()

	mongoClient, err := db.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB:", err)
	}

	rabbitConn, err := rabbitmq.ConnectRabbitMQ(cfg)
	if err != nil {
		log.Fatal("Cannot connect to RabbitMQ:", err)
	}

	r := gin.Default()

	// Pass the database client and RabbitMQ connection to the Gin context
	r.Use(func(c *gin.Context) {
		c.Set("mongo_client", mongoClient)
		c.Set("rabbit_conn", rabbitConn)
		c.Next()
	})

	api.SetupRoutes(r) 

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Cannot start the server:", err)
	}
}
