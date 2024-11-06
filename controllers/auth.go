package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"management_student/models"
	"net/http"
)

func Register(c *gin.Context, db *gorm.DB) {
	var registrationData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON body to registrationData struct
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create a new user with the provided information
	user := models.User{
		Email:    registrationData.Email,
		Password: registrationData.Password,
	}

	// Register the user in the database
	if err := models.RegisterUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"user": gin.H{
			"email": user.Email,
		},
	})
}

func Login(c *gin.Context, db *gorm.DB) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON body to loginData struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate user
	user, err := models.AuthenticateUser(db, loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Return the user's role and email
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"email": user.Email,
		},
	})
}
