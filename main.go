package main

import (
	"LESSON3_CRUD/api"
	"LESSON3_CRUD/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Connect to MongoDB
	database.ConnectMongoDB()
	// Set up routes
	// http.HandleFunc("/createProduct", api.CreateProductHandler)
	http.HandleFunc("/upd", api.CreateOrderHandler)
	// http.HandleFunc("/searchProducts", api.SearchProductsHandler)
	// http.HandleFunc("/getReport", api.GetReportHandler)
	
	// Start the serverve
	fmt.Println("server running on localhost 80:")
	log.Fatal(http.ListenAndServe(":8989", nil))
}

