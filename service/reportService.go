package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// ReportService handles generating reports for products and orders.
type ReportService struct {
	Collection *mongo.Collection
}

// GetReport generates a report of products with orders and total revenue.
func (service *ReportService) GetReport() ([]bson.M, error) {
	pipeline := mongo.Pipeline{
		// Match products with a quantity greater than or equal to 5
		{{
			Key:   "$match",
			Value: bson.D{{Key: "quantity", Value: bson.D{{Key: "$gte", Value: 5}}}},
		}},
		// Lookup orders for the products
		{{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "orders"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "productId"},
				{Key: "as", Value: "orders"},
			},
		}},
		// Unwind the orders array
		{{
			Key:   "$unwind",
			Value: "$orders",
		}},
		// Group by product name and calculate total orders and revenue
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$name"},
				{Key: "totalOrders", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "totalRevenue", Value: bson.D{{Key: "$sum", Value: "$orders.totalPrice"}}},
			},
		}},
		// Project the final output
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "product", Value: "$_id"},
				{Key: "totalOrders", Value: 1},
				{Key: "totalRevenue", Value: 1},
				{Key: "_id", Value: 0},
			},
		}},
		// Sort by total revenue in descending order
		{{
			Key:   "$sort",
			Value: bson.D{{Key: "totalRevenue", Value: -1}},
		}},
	}

	// Execute the aggregation pipeline
	cursor, err := service.Collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the results into a slice of bson.M
	var report []bson.M
	if err = cursor.All(context.Background(), &report); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return report, nil
}
