package pkg

import (
	"os"
)

// Config represents the application configuration.
type Config struct {
	Port       string
	DBHost     string
	APISecret  string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	return &Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		APISecret:  os.Getenv("API_SECRET"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}, nil
}
