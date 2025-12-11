package main

import (
	"log"

	"mpango-wa-cuddles/internal/config"
	"mpango-wa-cuddles/internal/db"
	"mpango-wa-cuddles/internal/handlers"
	"mpango-wa-cuddles/internal/models"
	"mpango-wa-cuddles/internal/router"
)

func main() {
	cfg := config.Load()

	conn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	if err := conn.AutoMigrate(&models.Location{}, &models.Ad{}, &models.User{}, &models.Ticket{}, &models.Saved{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	if err := handlers.Seed(conn); err != nil {
		log.Fatalf("seed failed: %v", err)
	}

	handlers.Init(conn, cfg.JWTSecret)

	r := router.New(cfg)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}


