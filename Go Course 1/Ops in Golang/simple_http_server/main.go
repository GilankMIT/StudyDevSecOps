package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for "/hello"
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hi! This is a sample page")
	})

	// Start the server on port 8080
	router.Run(":8080")
}