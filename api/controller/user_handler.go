package controller

import (
	"net/http"

	"api/model"
	"api/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userRepo repository.UserRepository // Interface for user operations

// InitUserHandler initializes the user repository
func InitUserHandler(db *mongo.Client) {
	userRepo = repository.NewUserRepository(db)
}

// SignUp handles user registration
func SignUp(c *gin.Context) {
	var signupRequest model.SignupRequestPayload
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// // Hash the password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	// 	return
	// }
	//signupRequest.Password = string(hashedPassword)

	// Check if the user already exists
	_, err := userRepo.FindUserByEmail(signupRequest.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check for existing user"})
		return
	}
	user := model.User{
		Name: signupRequest.Name,

		Email:        signupRequest.Email,
		Age:          signupRequest.Age,
		Password:     signupRequest.Password,
		MobileNumber: signupRequest.MobileNumber,
	}

	// Insert the new user
	if err := userRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
