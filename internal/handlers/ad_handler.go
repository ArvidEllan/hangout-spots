package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mpango-wa-cuddles/internal/models"
)

// ListAds returns partner ads sorted by weight (desc).
func ListAds(c *gin.Context) {
	var ads []models.Ad
	if err := db.Order("weight DESC").Find(&ads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load ads"})
		return
	}

	selected := weightedPick(ads)
	c.JSON(http.StatusOK, gin.H{"data": ads, "primary": selected, "count": len(ads)})
}

// CreateAd is a placeholder admin endpoint.
func CreateAd(c *gin.Context) {
	var ad models.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&ad).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create ad"})
		return
	}
	c.JSON(http.StatusCreated, ad)
}

// UpdateAd updates ad attributes.
func UpdateAd(c *gin.Context) {
	id := c.Param("id")
	var payload models.Ad
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Ad{}).Where("id = ?", id).Updates(payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update ad"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// TrackAdClick increments click count for tracking.
func TrackAdClick(c *gin.Context) {
	id := c.Param("id")
	if err := db.Model(&models.Ad{}).Where("id = ?", id).UpdateColumn("clicks", gorm.Expr("clicks + 1")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to track click"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "tracked"})
}

func weightedPick(ads []models.Ad) *models.Ad {
	if len(ads) == 0 {
		return nil
	}
	total := 0
	for _, ad := range ads {
		if ad.Weight > 0 {
			total += ad.Weight
		}
	}
	if total == 0 {
		return &ads[0]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := r.Intn(total) + 1
	acc := 0
	for i := range ads {
		acc += ads[i].Weight
		if roll <= acc {
			return &ads[i]
		}
	}
	return &ads[0]
}
