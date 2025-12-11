package models

import "github.com/google/uuid"

// Ad represents partner or affiliate promotions.
type Ad struct {
	ID       uuid.UUID `json:"id"`
	Type     string    `json:"type"`
	Link     string    `json:"link"`
	ImageURL string    `json:"image_url"`
	Weight   int       `json:"weight"`
}


