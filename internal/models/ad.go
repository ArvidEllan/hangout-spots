package models

import "github.com/google/uuid"

// Ad represents partner or affiliate promotions.
type Ad struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Type     string    `json:"type"`
	Link     string    `json:"link"`
	ImageURL string    `json:"image_url"`
	Weight   int       `json:"weight"`
	Clicks   int       `json:"clicks"`
}
