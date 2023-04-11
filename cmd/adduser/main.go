package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"log"
	"time"

	"github.com/SinisterSup/messageQueuing/internal/config"
	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInput struct {
	UserID    string  `json:"user_id"`
	Name      string  `json:"name"`
	Mobile    string  `json:"mobile"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	cfg := config.LoadConfig()

	mongoClient, err := db.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userCollection := db.GetCollection(mongoClient, "UserProduct-Messaging", "Users")

	// Read users from the JSON file
	file, err := os.Open("users.json")
	if err != nil {
		log.Fatalf("Failed to open users.json: %v", err)
	}
	defer file.Close()

	var users []UserInput
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		log.Fatalf("Failed to decode users data: %v", err)
	}

	for _, input := range users {
		// Check if the user with the given user_id already exists
		filter := bson.M{"_id": input.UserID}
		var existingUser models.User
		err := userCollection.FindOne(context.Background(), filter).Decode(&existingUser)
		
		if err == mongo.ErrNoDocuments {
			user := models.User{
				ID:        input.UserID,
				Name:      input.Name,
				Mobile:    input.Mobile,
				Latitude:  input.Latitude,
				Longitude: input.Longitude,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			_, err = userCollection.InsertOne(context.Background(), user)
			if err != nil {
				log.Fatalf("Failed to insert user: %v", err)
			}

			fmt.Printf("User '%s' with ID '%v' was successfully added.\n", user.Name, user.ID)
		} else if err == nil {
			// Update the existing user
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "name", Value: input.Name},
					{Key: "mobile", Value: input.Mobile},
					{Key: "latitude", Value: input.Latitude},
					{Key: "longitude", Value: input.Longitude},
					{Key: "updated_at", Value: time.Now()},
				}},
			}
			_, err := userCollection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				log.Fatalf("Failed to update user: %v", err)
			}

			fmt.Printf("User '%s' with ID '%v' was successfully updated.\n", input.Name, existingUser.ID)
		} else {
			log.Printf("Error occurred while searching for user: %v", err)
		}
	}
}
