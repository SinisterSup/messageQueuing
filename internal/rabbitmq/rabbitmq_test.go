package rabbitmq

import (
	"testing"

	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestConnectRabbitMQ(t *testing.T) {
	cfg := config.LoadConfig()
	conn, err := ConnectRabbitMQ(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	conn.Close()
}

func TestSendProductID(t *testing.T) {
	cfg := config.LoadConfig()
	conn, _ := ConnectRabbitMQ(cfg)
	defer conn.Close()

	productID := primitive.NewObjectID()
	err := SendProductID(conn, productID)
	assert.NoError(t, err)
}

func BenchmarkConnectRabbitMQ(b *testing.B) {
	cfg := config.LoadConfig()
	for i := 0; i < b.N; i++ {
		conn, _ := ConnectRabbitMQ(cfg)
		conn.Close()
	}
}

func BenchmarkSendProductID(b *testing.B) {
	cfg := config.LoadConfig()
	conn, _ := ConnectRabbitMQ(cfg)
	defer conn.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productID := primitive.NewObjectID()
		SendProductID(conn, productID)
	}
}
