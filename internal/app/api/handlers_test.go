package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/rabbitmq"
	"github.com/SinisterSup/messageQueuing/internal/app/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	api.SetupRoutes(r)
	return r
}


func TestCreateProduct(t *testing.T) {
	r := setupTestEnvironment(t)

	productData := api.InputAPI{
		UserID:            "testuser",
		ProductName:       "Test Product",
		ProductDescription: "This is a test product",
		ProductImages:     []string{"image1.jpg", "image2.jpg"},
		ProductPrice:      9.99,
	}

	payload, err := json.Marshal(productData)
	if err != nil {
		t.Errorf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/products", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = req

	w := httptest.NewRecorder()
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status code to be 201")
}
