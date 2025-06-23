package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppName    string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSchema   string
	AppPort    string
	JWTSecret  string
}

func LoadConfig() *Config {

	return &Config{
		AppName:    os.Getenv("APP_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSchema:   os.Getenv("DB_SCHEMA"),
		AppPort:    os.Getenv("PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
