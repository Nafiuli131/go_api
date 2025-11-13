package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nafiul/api_tutorial/controllers"
	"github.com/nafiul/api_tutorial/initializers"
	"github.com/nafiul/api_tutorial/migrate"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	migrate.MigrateDB(initializers.DB)
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)

	r := gin.Default()

	// User CRUD routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run()
}
