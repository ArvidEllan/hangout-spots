package models

import (
	"time"

	"github.com/google/uuid"
)

// Ticket tracks purchase attempts for venues.
type Ticket struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LocationID  uuid.UUID `json:"location_id" gorm:"type:uuid"`
	Price       int       `json:"price"`
	MpesaRef    string    `json:"mpesa_ref"`
	UserPhone   string    `json:"user_phone"`
	Status      string    `json:"status"` // pending, paid, failed
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}


