package handlers

import (
	"github.com/google/uuid"

	"mpango-wa-cuddles/internal/models"
)

var demoLocations = []models.Location{
	{
		ID:            uuid.New(),
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
		ID:            uuid.New(),
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
		ID:            uuid.New(),
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

var demoAds = []models.Ad{
	{
		ID:       uuid.New(),
		Type:     "uber",
		Link:     "https://m.uber.com/?promo=CUDDLE10",
		ImageURL: "https://images.unsplash.com/photo-1504198453319-5ce911bafcde",
		Weight:   5,
	},
	{
		ID:       uuid.New(),
		Type:     "giftshop",
		Link:     "https://flowers.ke/cuddles",
		ImageURL: "https://images.unsplash.com/photo-1501004318641-b39e6451bec6",
		Weight:   3,
	},
}

var demoTickets = []models.Ticket{}
