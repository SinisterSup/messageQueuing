package rabbitmq

import (
	"context"

	"github.com/SinisterSup/messageQueuing/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConnectRabbitMQ(cfg *config.Config) (*amqp.Connection, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.URI)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SendProductID(conn *amqp.Connection, productID primitive.ObjectID) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	q, err := channel.QueueDeclare("product_queue", false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = channel.PublishWithContext( // PublishWithContext is the new method in amqp091-go // `Publish` is deprecated
		context.Background(),
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(productID.Hex()), 
		},
	)

	if err != nil {
		return err
	}

	return nil
}
