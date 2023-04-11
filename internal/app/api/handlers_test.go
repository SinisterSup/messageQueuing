package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"context"
	// "time"
	"net/http/httptest"
	"testing"

	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/rabbitmq"
	"github.com/SinisterSup/messageQueuing/internal/app/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

func setupTestEnvironment(t *testing.T) *gin.Engine {
	cfg := config.LoadConfig()
	mockMongoClient, err := db.ConnectMongoDB(cfg)
	if err != nil {
		t.Error("Cannot connect to MongoDB:", err)
	}
	mockRabbitConn, err := rabbitmq.ConnectRabbitMQ(cfg)
	if err != nil {
		t.Error("Cannot connect to RabbitMQ:", err)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("mongo_client", mockMongoClient)
		c.Set("rabbit_conn", mockRabbitConn)
		c.Next()
	})
	route := r.Group("/api")
	{
		route.POST("/products", api.CreateProduct)
	}
	return r
}

// func createUser(client *mongo.Client, userID string) error {
// 	userCollection := db.GetCollection(client, "UserProduct-Messaging", "Users")
// 	_, err := userCollection.InsertOne(context.Background(), bson.M{"_id": userID, "created_at": time.Now()})
// 	return err
// }

func deleteUser(client *mongo.Client, userID string) error {
	userCollection := db.GetCollection(client, "UserProduct-Messaging", "Users")
	_, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": userID})
	return err
}

func TestCreateProduct(t *testing.T) {
	r := setupTestEnvironment(t)

	cfg := config.LoadConfig()
	mongoClient, err := db.ConnectMongoDB(cfg)
	if err != nil {
		t.Error("Couldn't connect to MongoDB:", err)
	}

	productData := api.InputAPI{
		UserID:            "u12345",
		ProductName:       "Test Product",
		ProductDescription: "This is a test product",
		ProductImages:     []string{"image1.jpg", "image2.jpg"},
		ProductPrice:      9.99,
	}
	defer deleteUser(mongoClient, productData.UserID)

	payload, err := json.Marshal(productData)
	if err != nil {
		t.Errorf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/products", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code to be 201 but got %d. Response body: %s", w.Code, w.Body.String())
	}
}
