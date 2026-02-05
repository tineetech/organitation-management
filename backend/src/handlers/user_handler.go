package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"organization-management-justin/backend/src/services"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func GetUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
