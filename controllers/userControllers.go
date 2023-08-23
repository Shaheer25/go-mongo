package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Shaheer25/go-mongo/initializers"
	"github.com/Shaheer25/go-mongo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {
	// Get the MongoDB database instance
	db, err := initializers.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle to the "users" collection
	collection := db.Collection("users")

	// Find all documents in the collection
	var users []models.User
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into a User struct
	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}


func CreateUser(c *gin.Context) {
	// Get the MongoDB database instance
	db, err := initializers.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle to the "users" collection
	collection := db.Collection("users")

	// Bind the request body to the User struct
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse user data",
		})
		return
	}

	// Insert the user document into the collection
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
