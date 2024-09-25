package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT         string
	DB_URL       string
	KAFKA_BROKER string
	KAFKA_TOPIC  string
}

func LoadConfig() *Config {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	envPath := filepath.Join(currentDir, "/.env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Could not load .env file from %s: %v", envPath, err)
	}

	return &Config{
		PORT:         os.Getenv("PORT"),
		DB_URL:       os.Getenv("DB_URL"),
		KAFKA_BROKER: os.Getenv("KAFKA_BROKER"),
		KAFKA_TOPIC:  os.Getenv("KAFKA_TOPIC"),
	}
}
