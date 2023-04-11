package api

import (
	"context"
	// "fmt"
	"net/http"
	"time"

	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/models"
	"github.com/SinisterSup/messageQueuing/internal/rabbitmq"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InputAPI struct {
	// ProductID 	      primitive.ObjectID  `json:"product_id" binding:"required"` // This is the ID of the product that will be sent to the queue
	UserID           string   `json:"user_id" binding:"required"`
	ProductName      string   `json:"product_name" binding:"required"`
	ProductDescription string `json:"product_description" binding:"required"`
	ProductImages     []string `json:"product_images" binding:"required"`
	ProductPrice      float64  `json:"product_price" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	var input InputAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := c.MustGet("mongo_client").(*mongo.Client)
	rabbitConn := c.MustGet("rabbit_conn").(*amqp.Connection)

	product := models.Product{
		ID: 			  primitive.NewObjectID(),
		Name:             input.ProductName,
		Description:      input.ProductDescription,
		Images:           input.ProductImages,
		Price:            input.ProductPrice,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	productCollection := db.GetCollection(client, "UserProduct-Messaging", "Products")
	_, err := productCollection.InsertOne(context.Background(), product)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the product"})
		return
	}

	userCollection := db.GetCollection(client, "UserProduct-Messaging", "Users")
	_, err = userCollection.InsertOne(context.Background(), bson.M{"_id": input.UserID, "created_at": time.Now()})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the user: user already exists or invalid user ID"})
		return
	}

	productID := product.ID
	err = rabbitmq.SendProductID(rabbitConn, productID) // Send the product ID to the message queue
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the product ID to the queue"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": bson.M{"_id": productID}})
}
