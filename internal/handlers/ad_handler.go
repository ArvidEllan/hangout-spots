package handlers

import (
	"net/http"
	"sort"

	"mpango-wa-cuddles/internal/models"

	"github.com/gin-gonic/gin"
)

// ListAds returns partner ads sorted by weight (desc).
func ListAds(c *gin.Context) {
	ads := make([]models.Ad, 0, len(demoAds))
	sorted := make([]adWithWeight, 0, len(demoAds))
	for _, ad := range demoAds {
		sorted = append(sorted, adWithWeight{Weight: ad.Weight, Payload: ad})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Weight > sorted[j].Weight
	})

	for _, a := range sorted {
		ads = append(ads, a.Payload)
	}

	c.JSON(http.StatusOK, gin.H{"data": ads, "count": len(ads)})
}

// CreateAd is a placeholder admin endpoint.
func CreateAd(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "stub - persist to DB in production"})
}

type adWithWeight struct {
	Weight  int
	Payload models.Ad
}
