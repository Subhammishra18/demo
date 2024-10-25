package main

import (
	"api/configuration"
	"api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := configuration.ConnectDB()

	// Initialize routes with the database connection
	routes.AuthRoutes(router, db)

	log.Fatal(router.Run(":8000"))
}
