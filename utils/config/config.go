package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT    string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
)

func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}

func LoadEnv() {
	primaryEnv := ".env"
	fallbackEnv := ".env-example"

	// Load primary .env
	err := godotenv.Overload(primaryEnv)
	if err != nil {
		log.Println(".env not found, trying .env-example as fallback")
		// Load fallback .env-example
		err = godotenv.Overload(fallbackEnv)
		if err != nil {
			log.Println("unable to load .env-example, running with default env")
		}
	}

	// It's possible set fallback to dev or local environments if needed
	PORT = GetEnv("PORT", "3333")
	DB_USER = GetEnv("DB_USER", "pismodbuser")
	DB_PASS = GetEnv("DB_PASS", "pismodbpassword")
	DB_HOST = GetEnv("DB_HOST", "localhost")
	DB_PORT = GetEnv("DB_PORT", "3306")
	DB_NAME = GetEnv("DB_NAME", "pismo-db")
}
