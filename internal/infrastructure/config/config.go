package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBSSL  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBSSL:  os.Getenv("DB_SSLMODE"),
	}
}
