package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"mpango-wa-cuddles/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListLocations returns curated places with optional filters.
func ListLocations(c *gin.Context) {
	cost := c.Query("cost") // under-200, 200-500, 500-1000
	area := c.Query("area") // Karen, CBD, etc.
	activity := c.Query("activity")

	results := make([]models.Location, 0, len(demoLocations))
	for _, loc := range demoLocations {
		if cost != "" && !withinCostBand(loc.CostPerPerson, cost) {
			continue
		}
		if area != "" && !strings.EqualFold(loc.Area, area) {
			continue
		}
		if activity != "" && !hasActivity(loc.Activities, activity) {
			continue
		}
		results = append(results, loc)
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

	for _, loc := range demoLocations {
		if loc.ID == id {
			c.JSON(http.StatusOK, loc)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
}

// CreateLocation is a placeholder admin endpoint.
func CreateLocation(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "stub - persist to DB in production"})
}

// UpdateLocation is a placeholder admin endpoint.
func UpdateLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "stub - update DB row"})
}

// DeleteLocation is a placeholder admin endpoint.
func DeleteLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "stub - delete DB row"})
}

func withinCostBand(cost int, band string) bool {
	switch strings.ToLower(band) {
	case "under-200":
		return cost < 200
	case "200-500":
		return cost >= 200 && cost <= 500
	case "500-1000":
		return cost > 500 && cost <= 1000
	default:
		// allow numeric single value filter
		if v, err := strconv.Atoi(band); err == nil {
			return cost <= v
		}
		return true
	}
}

func hasActivity(list []string, target string) bool {
	for _, a := range list {
		if strings.EqualFold(a, target) {
			return true
		}
	}
	return false
}
