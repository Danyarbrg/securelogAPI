package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port		string
	DatabaseURL string
}

func LoadConfig() Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	cfg := Config {
		Port:		os.Getenv("PORT"),
		DatabaseURL:	os.Getenv("DATABASE_URL"),
	}

	return cfg
}
