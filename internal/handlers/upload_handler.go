package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadImage is a placeholder that should be replaced with S3 upload logic.
func UploadImage(c *gin.Context) {
	// In production: accept multipart file, upload to S3, return URL.
	c.JSON(http.StatusOK, gin.H{
		"image_url": "https://placehold.co/600x400.png",
		"note":      "Replace with S3 upload implementation",
	})
}

