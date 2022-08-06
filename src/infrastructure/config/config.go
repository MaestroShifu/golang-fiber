package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	PROD = "production"
	DEV  = "development"
)

type Config struct {
	NAME_APP          string `json:"NAME_APP"`
	APP_ENV           string `json:"APP_ENV"`
	DB_HOST           string `json:"DB_HOST"`
	DB_USER           string `json:"DB_USER"`
	POSTGRES_PASSWORD string `json:"POSTGRES_PASSWORD"`
	DB_NAME           string `json:"DB_NAME"`
}

func getPathEnv() string {
	abs, err := filepath.Abs(".env")
	if err == nil {
		log.Println("Absolute path is: ", abs)
	}
	return abs
}

func GetConfigSystem() Config {
	APP_ENV := os.Getenv("APP_ENV")

	if APP_ENV != PROD {
		err := godotenv.Load(getPathEnv())
		if err != nil {
			log.Fatal("Error loading env")
		}
	}

	return Config{
		NAME_APP:          os.Getenv("NAME_APP"),
		APP_ENV:           APP_ENV,
		DB_HOST:           os.Getenv("DB_HOST"),
		DB_USER:           os.Getenv("DB_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		DB_NAME:           os.Getenv("DB_NAME"),
	}
}
