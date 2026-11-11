package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func getEnvOrError(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("environment variable %s is required but not set", key)
	}

	return value, nil
}

func Load() (*Config, error) {
	port, err := getEnvOrError("PORT")
	if err != nil {
		return nil, err
	}

	dbAddress, err := getEnvOrError("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbAddress,
	}, nil
}
