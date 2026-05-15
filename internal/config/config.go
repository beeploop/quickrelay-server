package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	instance *Config
)

type Config struct {
	PORT string
}

func Load() *Config {
	if instance != nil {
		return instance
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env: %s\n", err.Error())
	}

	config := &Config{
		PORT: mustGetEnv("PORT", ":3000"),
	}

	return config
}

func mustGetEnv(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	return value
}
