package handlers

import "github.com/gin-gonic/gin"

// Health is a simple readiness probe.
func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}


