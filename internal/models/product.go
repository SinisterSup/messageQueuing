package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                      primitive.ObjectID    `bson:"_id,omitempty"`
	Name                    string    `bson:"product_name"`
	Description             string    `bson:"product_description"`
	Images                  []string  `bson:"product_images"`
	Price                   float64   `bson:"product_price"`
	CompressedProductImages []string  `bson:"compressed_product_images"`
	CreatedAt               time.Time `bson:"created_at"`
	UpdatedAt               time.Time `bson:"updated_at"`
}
