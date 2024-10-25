package routes

import (
	"api/controller"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthRoutes sets up authentication-related routes
func AuthRoutes(router *gin.Engine, db *mongo.Client) {
	controller.InitUserHandler(db) // Call to initialize user-related handlers
	auth := router.Group("/auth")  // Create a new route group for authentication
	{
		auth.POST("/signup", controller.SignUp) // Route for user signup
		auth.POST("/login", controller.Login)   // Route for user login
	}
}
