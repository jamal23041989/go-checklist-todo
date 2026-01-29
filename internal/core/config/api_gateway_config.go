package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type APIGatewayConfig struct {
	Port          string        `env:"HTTP_PORT"`
	Timeout       time.Duration `env:"HTTP_TIMEOUT"`
	DBCoreAddress string        `env:"DB_CORE_ADDRESS"`
	LogLevel      string        `env:"LOG_LEVEL"`
}

func LoadAPIGatewayConfig() (*APIGatewayConfig, error) {
	var cfg APIGatewayConfig

	cfg.Port = getEnvOrDefault("HTTP_PORT", "8080")
	cfg.Timeout = cfg.parseDurationOrDefault("HTTP_TIMEOUT", "30s")
	cfg.DBCoreAddress = getEnvOrDefault("DB_CORE_ADDRESS", "localhost:50051")
	cfg.LogLevel = getEnvOrDefault("LOG_LEVEL", "info")

	if err := validateAPIGatewayConfig(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (g *APIGatewayConfig) parseDurationOrDefault(key, defaultValue string) time.Duration {
	envValue := os.Getenv(key)
	if envValue == "" {
		// Попробуем парсить defaultValue в time.Duration
		duration, err := time.ParseDuration(defaultValue)
		if err != nil {
			return 30 * time.Second
		}
		return duration
	}

	duration, err := time.ParseDuration(envValue)
	if err != nil {
		return 30 * time.Second
	}

	return duration
}

func validateAPIGatewayConfig(g APIGatewayConfig) error {
	if g.Port == "" {
		return errors.New("invalid API gateway port")
	}

	if _, err := strconv.Atoi(g.Port); err != nil {
		return errors.New("port must be a number")
	}

	if g.DBCoreAddress == "" {
		return errors.New("invalid DB core address")
	}

	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}
	return defaultValue
}
