package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a minimal email-only account.
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"-"` // hashed
	CreatedAt time.Time `json:"created_at"`
}
