package handlers

import (
	"log"

	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	jwtSecret []byte
)

// Init wires shared dependencies for handlers.
func Init(database *gorm.DB, secret string) {
	db = database
	jwtSecret = []byte(secret)
	if db == nil {
		log.Fatal("handlers.Init: database is nil")
	}
}


