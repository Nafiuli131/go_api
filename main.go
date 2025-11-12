package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nafiul/api_tutorial/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello this is from Nafiul Islam",
		})
	})
	r.Run()
}
