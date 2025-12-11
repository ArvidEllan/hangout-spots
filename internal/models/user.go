package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a minimal email-only account.
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // hashed
	CreatedAt time.Time `json:"created_at"`
}


