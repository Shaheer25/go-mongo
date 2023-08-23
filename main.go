package main

import (


	"github.com/Shaheer25/go-mongo/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register the route to get all users
	r.GET("/getusers", controllers.GetUsers)

	// Register the route to create a new user
	r.POST("/createusers", controllers.CreateUser)

	// Run the server
	r.Run(":3000")
}
