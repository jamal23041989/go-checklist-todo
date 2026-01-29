// Package config provides configuration management for all microservices.
// Each service has its own configuration structure with environment variable loading
// and validation. Configurations are loaded at application startup.
package config

import (
	"errors"
	"strconv"
	"time"

	"github.com/jamal23041989/go-checklist-todo/internal/core/tools"
)

// APIGatewayConfig holds configuration for the API Gateway microservice.
// This service handles HTTP requests from users and forwards them to DB Core service.
//
// Environment Variables:
//   - HTTP_PORT: Port for HTTP server (default: "8080")
//   - HTTP_TIMEOUT: Request timeout duration (default: "30s")
//   - DB_CORE_ADDRESS: gRPC address of DB Core service (default: "localhost:50051")
//   - LOG_LEVEL: Logging level - debug, info, warn, error (default: "info")
type APIGatewayConfig struct {
	Port          string        `env:"HTTP_PORT"`       // HTTP server port
	Timeout       time.Duration `env:"HTTP_TIMEOUT"`    // Request timeout
	DBCoreAddress string        `env:"DB_CORE_ADDRESS"` // gRPC service address
	LogLevel      string        `env:"LOG_LEVEL"`       // Logging level
}

// LoadAPIGatewayConfig loads API Gateway configuration from environment variables.
// Uses default values for missing environment variables and validates required fields.
//
// Returns:
//   - *APIGatewayConfig: Loaded configuration
//   - error: Validation error if required fields are missing or invalid
//
// Example:
//
//	cfg, err := LoadAPIGatewayConfig()
//	if err != nil {
//	    log.Fatal("Failed to load config:", err)
//	}
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

// validateAPIGatewayConfig validates API Gateway configuration fields.
// Ensures all required fields are present and valid.
//
// Parameters:
//   - g: APIGatewayConfig instance to validate
//
// Returns:
//   - error: Validation error if any field is invalid
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
