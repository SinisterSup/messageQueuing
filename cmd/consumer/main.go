package main

import (
	"fmt"
	"log"

	"github.com/SinisterSup/messageQueuing/internal/app/consumer"
	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/rabbitmq"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	channel, err := rabbitConn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer channel.Close()

	q, err := channel.QueueDeclare("product_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare the queue:", err)
	}

	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to consume messages from the queue:", err)
	}

	forever := make(chan bool)

	go func() {
		for message := range msgs {
			productID := string(message.Body)
			fmt.Printf("Received Product ID: %s\n", productID)

			productObjectID, err := primitive.ObjectIDFromHex(productID)
			if err != nil {
				log.Printf("Failed to convert productID to ObjectID: %s\n", productID)
				continue
			}

			err = consumer.DownloadCompressAndStoreImages(mongoClient, productObjectID) // Download, compress and store images for the product
			if err != nil {
				log.Printf("Failed to download, compress and store images for Product ID: %s\n", productID)
			}
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
