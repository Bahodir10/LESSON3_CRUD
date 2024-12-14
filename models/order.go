package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID string             `bson:"productId"`
	Quantity  int                `bson:"quantity"`
	TotalPrice float64           `bson:"totalPrice"`
}
