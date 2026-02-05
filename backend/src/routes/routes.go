package routes

import (
	"github.com/gin-gonic/gin"
	"organization-management-justin/backend/src/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// api.GET("/health", handlers.HealthCheck)
	api.GET("/users", handlers.GetUsers)
}
