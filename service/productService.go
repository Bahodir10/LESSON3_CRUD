package service

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"LESSON3_CRUD/models"
	"fmt"
)

// ProductService handles operations related to Products.
type ProductService struct {
	Collection *mongo.Collection
}

// CreateProduct adds a new product to the database.
func (service *ProductService) CreateProduct(product models.Product) error {
	_, err := service.Collection.InsertOne(context.Background(), product)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// GetProducts fetches a paginated list of products.
func (service *ProductService) GetProducts(page, limit int) ([]models.Product, error) {
	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := service.Collection.Find(context.Background(), bson.D{}, options)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	if err = cursor.All(context.Background(), &products); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return products, nil
}

// SearchProducts searches for products by name.
func (service *ProductService) SearchProducts(searchTerm string) ([]models.Product, error) {
	filter := bson.D{
		{Key: "name", Value: bson.D{
			{Key: "$regex", Value: searchTerm},
			{Key: "$options", Value: "i"},
		}},
	}
	

	cursor, err := service.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	if err = cursor.All(context.Background(), &products); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return products, nil
}


// UpdateProduct updates a product's information by ID.
func (service *ProductService) UpdateProduct(id primitive.ObjectID, updatedProduct models.Product) error {
	filter := bson.D{{Key:"_id", Value:id}}
	update := bson.D{
		{Key:"$set", Value:bson.D{
			{Key:"name", Value:updatedProduct.Name},
			{Key:"price", Value:updatedProduct.Price},
			{Key:"quantity", Value:updatedProduct.Quantity},
		}},
	}

	_, err := service.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// DeleteProduct deletes a product by ID.
func (service *ProductService) DeleteProduct(id primitive.ObjectID) error {
	_, err := service.Collection.DeleteOne(context.Background(), bson.D{{Key:"_id", Value:id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// GetProductByID retrieves a product by its ID.
func (service *ProductService) GetProductByID(id primitive.ObjectID) (models.Product, error) {
	var product models.Product
	filter := bson.D{{Key: "_id", Value: id}}

	// Find the product by ID
	err := service.Collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If no document is found, return an error indicating the product was not found
			return product, fmt.Errorf("product not found")
		}
		log.Fatal(err)
		return product, err
	}

	return product, nil
}
