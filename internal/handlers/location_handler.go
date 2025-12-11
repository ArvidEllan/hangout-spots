package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"mpango-wa-cuddles/internal/models"
)

// ListLocations returns curated places with optional filters.
func ListLocations(c *gin.Context) {
	cost := c.Query("cost") // under-200, 200-500, 500-1000
	area := c.Query("area") // Karen, CBD, etc.
	activity := c.Query("activity")

	query := db.Model(&models.Location{})

	if cost != "" {
		cond, args := costCondition(cost)
		if cond != "" {
			query = query.Where(cond, args...)
		}
	}
	if area != "" {
		query = query.Where("LOWER(area) = LOWER(?)", area)
	}
	if activity != "" {
		query = query.Where("? = ANY(activities)", activity)
	}

	results := []models.Location{}
	if err := query.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch locations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "count": len(results)})
}

// GetLocation retrieves a single location by id.
func GetLocation(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var loc models.Location
	if err := db.First(&loc, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}

	c.JSON(http.StatusOK, loc)
}

// CreateLocation is a placeholder admin endpoint.
func CreateLocation(c *gin.Context) {
	var loc models.Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&loc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create location"})
		return
	}

	c.JSON(http.StatusCreated, loc)
}

// UpdateLocation is a placeholder admin endpoint.
func UpdateLocation(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var body models.Location
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Location{}).Where("id = ?", id).Updates(body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteLocation is a placeholder admin endpoint.
func DeleteLocation(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := db.Delete(&models.Location{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func costCondition(band string) (string, []interface{}) {
	switch strings.ToLower(band) {
	case "under-200":
		return "cost_per_person < ?", []interface{}{200}
	case "200-500":
		return "cost_per_person BETWEEN ? AND ?", []interface{}{200, 500}
	case "500-1000":
		return "cost_per_person > ? AND cost_per_person <= ?", []interface{}{500, 1000}
	default:
		if v, err := strconv.Atoi(band); err == nil {
			return "cost_per_person <= ?", []interface{}{v}
		}
		return "", nil
	}
}
