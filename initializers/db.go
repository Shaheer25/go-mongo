package initializers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Database, error) {
	// MongoDB connection string
	dsn := "mongodb://localhost:27017"

	// Create a new MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Get a handle to the "mydatabase" database
	db := client.Database("mongo")

	return db, nil
}
