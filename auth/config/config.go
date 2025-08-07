package config

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	DBUrl     string
	JWTSecret string
	JWTExpire time.Duration
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		DBUrl:     os.Getenv("DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpire: time.Hour * 24,
	}
}
