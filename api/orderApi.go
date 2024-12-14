package api

import (
	"LESSON3_CRUD/database"
	"LESSON3_CRUD/models"
	"LESSON3_CRUD/service"
	"encoding/json"
	"net/http"
	"strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrderHandler creates a new order.
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	orderService := service.OrderService{Collection: database.OrdersCollection}

	err := orderService.CreateOrder(order)
	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Order created successfully"))
}

// GetOrdersHandler fetches orders with pagination.
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	orderService := service.OrderService{Collection: database.OrdersCollection}

	orders, err := orderService.GetOrders(page, limit)
	if err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}

	// Return orders as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// SearchOrdersHandler searches orders based on a search term.
func SearchOrdersHandler(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("search")
	orderService := service.OrderService{Collection: database.OrdersCollection}

	orders, err := orderService.SearchOrders(searchTerm)
	if err != nil {
		http.Error(w, "Error searching for orders", http.StatusInternalServerError)
		return
	}

	// Return orders as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// UpdateOrderHandler updates an order based on its ID.
func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var updatedOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	orderService := service.OrderService{Collection: database.OrdersCollection}
	err = orderService.UpdateOrder(id, updatedOrder)
	if err != nil {
		http.Error(w, "Error updating order", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Order updated successfully"))
}

// DeleteOrderHandler deletes an order by its ID.
func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderService := service.OrderService{Collection: database.OrdersCollection}
	err = orderService.DeleteOrder(id)
	if err != nil {
		http.Error(w, "Error deleting order", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Order deleted successfully"))
}

func GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the order ID from URL parameters
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid order ID format", http.StatusBadRequest)
		return
	}

	// Fetch the order using OrderService
	orderService := service.OrderService{Collection: database.OrdersCollection}
	order, err := orderService.GetOrderByID(objectID)
	if err != nil {
		http.Error(w, "Error retrieving order", http.StatusInternalServerError)
		return
	}

	// Respond with the order data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

