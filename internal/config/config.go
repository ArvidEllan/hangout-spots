package config

import "os"

// Config holds minimal environment configuration.
type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
	MpesaKey    string
	MpesaSecret string
	MpesaShort  string
	MpesaPass   string
	MpesaCallback string
}

// Load returns configuration from environment with sensible defaults.
func Load() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        firstNonEmpty(os.Getenv("PORT"), "8080"),
		JWTSecret:   firstNonEmpty(os.Getenv("JWT_SECRET"), "dev-secret"),
		MpesaKey:    os.Getenv("MPESA_CONSUMER_KEY"),
		MpesaSecret: os.Getenv("MPESA_CONSUMER_SECRET"),
		MpesaShort:  os.Getenv("MPESA_SHORTCODE"),
		MpesaPass:   os.Getenv("MPESA_PASSKEY"),
		MpesaCallback: os.Getenv("MPESA_CALLBACK_URL"),
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


