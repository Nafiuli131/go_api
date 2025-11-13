package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nafiul/api_tutorial/initializers"
	"github.com/nafiul/api_tutorial/models"
)

// Create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		var checkUser models.User 
		initializers.DB.Raw("select * from users where email=? ",user.Email).Scan(&checkUser)
		if checkUser.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Found Duplicate Email"})
			return
		}
	initializers.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update user by ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Age != 0 {
		user.Age = input.Age
	}

	initializers.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Delete user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
