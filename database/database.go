package database

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref" 
	"time"
)

var MongoClient *mongo.Client
var ProductsCollection *mongo.Collection
var OrdersCollection *mongo.Collection

func ConnectMongoDB() {
	mongoURI := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create and connect the client in one step
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Test the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB successfully!")
}
