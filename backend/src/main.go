package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"organization-management-justin/backend/src/routes"
)

func main() {
	r := gin.Default()

	// register routes
	routes.RegisterRoutes(r)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
