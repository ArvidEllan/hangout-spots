package handlers

import (
	"errors"

	"gorm.io/gorm"

	"mpango-wa-cuddles/internal/models"
)

var seedLocations = []models.Location{
	{
		Name:          "Oloolua Nature Trail",
		Area:          "Karen",
		CostPerPerson: 300,
		Description:   "Shaded boardwalks, waterfall, and caves for a calm stroll.",
		Activities:    []string{"Nature", "Hiking", "Picnic"},
		Notes:         "Carry water and light snacks.",
		TransportTip:  "Matatu to Karen + boda to the gate.",
		ImageURL:      "https://images.unsplash.com/photo-1500530855697-b586d89ba3ee",
		Ticketing:     true,
	},
	{
		Name:          "Paradise Lost",
		Area:          "Kiambu",
		CostPerPerson: 400,
		Description:   "Coffee farm lake with boat rides and caves.",
		Activities:    []string{"Adventure", "Boat", "Nature"},
		Notes:         "Boats charged separately. Great for golden hour photos.",
		TransportTip:  "Thika road matatu to Kiambu town then boda.",
		ImageURL:      "https://images.unsplash.com/photo-1507525428034-b723cf961d3e",
		Ticketing:     true,
	},
	{
		Name:          "Karura Forest Picnic",
		Area:          "Gigiri",
		CostPerPerson: 200,
		Description:   "Tall trees, cycling, and picnic lawns close to the city.",
		Activities:    []string{"Picnic", "Cycling", "Nature"},
		Notes:         "Pack mosquito repellent. Bikes for hire inside.",
		TransportTip:  "Matatu to Limuru road; short walk to gate A.",
		ImageURL:      "https://images.unsplash.com/photo-1469474968028-56623f02e42e",
		Ticketing:     false,
	},
}

var seedAds = []models.Ad{
	{
		Type:     "uber",
		Link:     "https://m.uber.com/?promo=CUDDLE10",
		ImageURL: "https://images.unsplash.com/photo-1504198453319-5ce911bafcde",
		Weight:   5,
	},
	{
		Type:     "giftshop",
		Link:     "https://flowers.ke/cuddles",
		ImageURL: "https://images.unsplash.com/photo-1501004318641-b39e6451bec6",
		Weight:   3,
	},
}

// Seed inserts demo data if tables are empty.
func Seed(db *gorm.DB) error {
	if db == nil {
		return errors.New("seed: db is nil")
	}

	var count int64

	db.Model(&models.Location{}).Count(&count)
	if count == 0 {
		if err := db.Create(&seedLocations).Error; err != nil {
			return err
		}
	}

	db.Model(&models.Ad{}).Count(&count)
	if count == 0 {
		if err := db.Create(&seedAds).Error; err != nil {
			return err
		}
	}
	return nil
}
