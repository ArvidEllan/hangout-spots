package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect opens a GORM connection to PostgreSQL.
func Connect(url string) (*gorm.DB, error) {
	if url == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}

	db, err := gorm.Open(postgres.Open(url), cfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return db, nil
}
