package controller

import (
	"api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// Import bcrypt package
)

func Login(c *gin.Context) {
	var loginRequest model.LoginRequestPayload
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	log.Printf("Login attempt for mobile number: %v", loginRequest.MobileNumber)

	// Find the user by mobile number
	user, err := userRepo.FindUserByMobile(loginRequest.MobileNumber)
	if err != nil {
		log.Printf("Login error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid mobile number or password"})
		return
	}

	log.Printf("User found with email: %v", user.Email)

	// // Check the password
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
	// 	log.Printf("Password comparison error for user %v: %v", user.Email, err)
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid mobile number or password"})
	// 	return
	// }
	if user.Password == loginRequest.Password {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		return
	}

	log.Printf("Login successful for user: %v", user.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
