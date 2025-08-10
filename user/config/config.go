package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBUrl string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		DBUrl: os.Getenv("DB_URL"),
	}
}
