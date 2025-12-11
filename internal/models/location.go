package models

import "github.com/google/uuid"

// Location represents a curated place for couples.
type Location struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Area          string    `json:"area"`
	CostPerPerson int       `json:"cost_per_person"`
	Description   string    `json:"description"`
	Activities    []string  `json:"activities"`
	Notes         string    `json:"notes"`
	TransportTip  string    `json:"transport_tip"`
	ImageURL      string    `json:"img_url"`
	Ticketing     bool      `json:"ticketing"`
}


