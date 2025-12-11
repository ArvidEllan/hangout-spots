package models

import "github.com/google/uuid"

// Saved represents a user's saved (cuddle list) entry.
type Saved struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;index"`
	LocationID uuid.UUID `json:"location_id" gorm:"type:uuid;index"`
}

