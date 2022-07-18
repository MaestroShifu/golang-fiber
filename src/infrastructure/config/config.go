package config

import (
	"fmt"
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
	NAME_APP string `json:"NAME_APP"`
	APP_ENV  string `json:"APP_ENV"`
}

func getPathEnv() string {
	abs, err := filepath.Abs(".env")
	if err == nil {
		fmt.Println("Absolute path is:", abs)
	}
	return abs
}

func GetConfigSystem() Config {
	APP_ENV := os.Getenv("APP_ENV")

	if APP_ENV == DEV {
		err := godotenv.Load(getPathEnv())

		if err != nil {
			log.Fatal("Error loading env")
		}
	}

	return Config{
		NAME_APP: os.Getenv("NAME_APP"),
		APP_ENV:  APP_ENV,
	}
}
