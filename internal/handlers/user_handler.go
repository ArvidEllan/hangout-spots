package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mpango-wa-cuddles/internal/models"
)

// Register is a placeholder for email-only signup.
func Register(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "stub - implement JWT auth"})
}

// Login is a placeholder for issuing JWTs.
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": "mock-token"})
}

// SavedPlaces returns a mock cuddle list.
func SavedPlaces(c *gin.Context) {
	if len(demoLocations) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": demoLocations[:1], "count": 1})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": []models.Location{}, "count": 0})
	}
}


