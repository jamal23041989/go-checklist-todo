package config

import (
	"errors"
	"strconv"
	"time"

	"github.com/jamal23041989/go-checklist-todo/internal/core/tools"
)

type APIGatewayConfig struct {
	Port          string        `env:"HTTP_PORT"`
	Timeout       time.Duration `env:"HTTP_TIMEOUT"`
	DBCoreAddress string        `env:"DB_CORE_ADDRESS"`
	LogLevel      string        `env:"LOG_LEVEL"`
}

func LoadAPIGatewayConfig() (*APIGatewayConfig, error) {
	apiGatewayConfig := &APIGatewayConfig{
		Port:          tools.GetEnvOrDefault("HTTP_PORT", "8080"),
		Timeout:       tools.ParseDurationOrDefault("HTTP_TIMEOUT", "30s"),
		DBCoreAddress: tools.GetEnvOrDefault("DB_CORE_ADDRESS", "localhost:50051"),
		LogLevel:      tools.GetEnvOrDefault("LOG_LEVEL", "info"),
	}

	if err := validateAPIGatewayConfig(*apiGatewayConfig); err != nil {
		return nil, err
	}

	return apiGatewayConfig, nil
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
