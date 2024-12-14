package service

import (
	"LESSON3_CRUD/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" // Ensure the import of options
	"fmt"
)

type OrderService struct {
	Collection *mongo.Collection
}

// CreateOrder creates a new order in the database.
func (service *OrderService) CreateOrder(order models.Order) error {
	_, err := service.Collection.InsertOne(context.Background(), order)
	return err
}

// GetOrders fetches orders with pagination.
func (service *OrderService) GetOrders(page, limit int) ([]models.Order, error) {
	// Convert limit and skip to int64
	skip := int64((page - 1) * limit)
	limitInt64 := int64(limit)

	// Fetch orders with pagination
	cursor, err := service.Collection.Find(context.Background(), bson.D{}, &options.FindOptions{
		Skip:  &skip,       // Skip as int64
		Limit: &limitInt64, // Limit as int64
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []models.Order
	if err = cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// SearchOrders searches orders based on a search term.
func (service *OrderService) SearchOrders(searchTerm string) ([]models.Order, error) {
	cursor, err := service.Collection.Find(context.Background(), bson.D{
		{Key: "productId", Value: bson.D{{Key: "$regex", Value: searchTerm}}},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []models.Order
	if err = cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder updates an order based on its ID.
func (service *OrderService) UpdateOrder(id primitive.ObjectID, updatedOrder models.Order) error {
	_, err := service.Collection.UpdateOne(context.Background(), bson.D{
		{Key: "_id", Value: id},
	}, bson.D{
		{Key: "$set", Value: updatedOrder},
	})
	return err
}

// DeleteOrder deletes an order by its ID.
func (service *OrderService) DeleteOrder(id primitive.ObjectID) error {
	_, err := service.Collection.DeleteOne(context.Background(), bson.D{
		{Key: "_id", Value: id},
	})
	return err
}

// GetOrderByID retrieves an order by its ID.
func (service *OrderService) GetOrderByID(id primitive.ObjectID) (models.Order, error) {
	var order models.Order
	filter := bson.D{{Key: "_id", Value: id}}

	// Find the order by ID
	err := service.Collection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If no document is found, return an error indicating the order was not found
			return order, fmt.Errorf("order not found")
		}
		return order, err
	}

	return order, nil
}

