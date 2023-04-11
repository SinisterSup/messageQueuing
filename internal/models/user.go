package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Mobile    string    `bson:"mobile"`
	Latitude  float64   `bson:"latitude"`
	Longitude float64   `bson:"longitude"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
