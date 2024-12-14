package api

import (
	"LESSON3_CRUD/database"
	"LESSON3_CRUD/models"
	"LESSON3_CRUD/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Decode the request body into the product struct
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	productService := service.ProductService{Collection: database.ProductsCollection}

	// Call the CreateProduct method
	err = productService.CreateProduct(product)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product created successfully"))
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    productService := service.ProductService{Collection: database.ProductsCollection}

    products, err := productService.GetProducts(page, limit)
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }

    for _, product := range products {
        fmt.Fprintf(w, "%v\n", product)
    }
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("search")
    productService := service.ProductService{Collection: database.ProductsCollection}
    products, err := productService.SearchProducts(searchTerm)
    if err != nil {
        http.Error(w, "Error searching for products", http.StatusInternalServerError)
        return
    }

    for _, product := range products {
        fmt.Fprintf(w, "%v\n", product)
    }
}

func GetReportHandler(w http.ResponseWriter, r *http.Request) {
    reportService := service.ReportService{Collection: database.ProductsCollection}

    report, err := reportService.GetReport()
    if err != nil {
        http.Error(w, "Error generating report", http.StatusInternalServerError)
        return
    }

    for _, item := range report {
        fmt.Fprintf(w, "%v\n", item)
    }
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from URL query parameter
	id := r.URL.Query().Get("id")
	productService := service.ProductService{Collection: database.ProductsCollection}

	// Convert id to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Parse the incoming request body to get updated product details
	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service method to update the product
	err = productService.UpdateProduct(objectID, updatedProduct)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.Write([]byte("Product updated successfully"))
}



func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	productService := service.ProductService{Collection: database.ProductsCollection}

	// Convert id to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = productService.DeleteProduct(objectID)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product deleted successfully"))
}


func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from URL parameters
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	// Fetch the product using ProductService
	productService := service.ProductService{Collection: database.ProductsCollection}
	product, err := productService.GetProductByID(objectID)
	if err != nil {
		http.Error(w, "Error retrieving product", http.StatusInternalServerError)
		return
	}

	// Respond with the product data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

