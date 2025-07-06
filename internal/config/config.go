package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
}

// Load reads environment variables from .env and returns a Config struct
func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DiscordToken: getEnv("DISCORD_TOKEN"),
		DBHost:       getEnv("DB_HOST"),
		DBPort:       getEnv("DB_PORT"),
		DBUser:       getEnv("DB_USER"),
		DBPassword:   getEnv("DB_PASSWORD"),
		DBName:       getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
