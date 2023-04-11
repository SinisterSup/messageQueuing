package db

import (
	"context"
	"testing"
	// "time"

	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/stretchr/testify/assert"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnectMongoDB(t *testing.T) {
	cfg := config.LoadConfig()
	client, err := ConnectMongoDB(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	err = client.Disconnect(context.Background())
	assert.NoError(t, err)
}

func TestGetCollection(t *testing.T) {
	cfg := config.LoadConfig()
	client, _ := ConnectMongoDB(cfg)
	defer client.Disconnect(context.Background())

	collection := GetCollection(client, "test_db", "test_collection")
	assert.NotNil(t, collection)
	assert.IsType(t, &mongo.Collection{}, collection)

	collection.Drop(context.Background())
}

func BenchmarkConnectMongoDB(b *testing.B) {
	cfg := config.LoadConfig()
	for i := 0; i < b.N; i++ {
		client, _ := ConnectMongoDB(cfg)
		client.Disconnect(context.Background())
	}
}
