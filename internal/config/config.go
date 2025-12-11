package config

import "os"

// Config holds minimal environment configuration.
type Config struct {
	DatabaseURL string
	Port        string
}

// Load returns configuration from environment with sensible defaults.
func Load() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        firstNonEmpty(os.Getenv("PORT"), "8080"),
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}


