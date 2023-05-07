package config

import (
	"github.com/joho/godotenv"
	"log"
)

func InitializeEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file")
	}
}
