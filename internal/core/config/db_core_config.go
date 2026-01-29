// Package config provides configuration management for all microservices.
// Each service has its own configuration structure with environment variable loading
// and validation. Configurations are loaded at application startup.
package config

import (
	"errors"

	"github.com/jamal23041989/go-checklist-todo/internal/core/tools"
)

// PostgresConfig holds configuration for PostgreSQL database connection.
// Used by DB Core service for persistent task storage.
//
// Environment Variables:
//   - POSTGRES_HOST: PostgreSQL server host (default: "localhost")
//   - POSTGRES_PORT: PostgreSQL server port (default: "5432")
//   - POSTGRES_USER: Database username (required)
//   - POSTGRES_PASSWORD: Database password (required)
//   - POSTGRES_DB: Database name (required)
//   - POSTGRES_SSLMODE: SSL mode - disable, require, verify-ca, verify-full (default: "disable")
type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`     // Database server host
	Port     string `env:"POSTGRES_PORT"`     // Database server port
	User     string `env:"POSTGRES_USER"`     // Database username
	Password string `env:"POSTGRES_PASSWORD"` // Database password
	Name     string `env:"POSTGRES_DB"`       // Database name
	SSLMode  string `env:"POSTGRES_SSLMODE"`  // SSL connection mode
}

// validatePostgresConfig validates PostgreSQL configuration fields.
// Ensures all required fields are present and SSL mode is valid.
//
// Parameters:
//   - p: PostgresConfig instance to validate
//
// Returns:
//   - error: Validation error if any field is invalid
func validatePostgresConfig(p PostgresConfig) error {
	if p.Host == "" {
		return errors.New("host required")
	}

	if p.Port == "" {
		return errors.New("port required")
	}

	if len([]rune(p.Password)) == 0 {
		return errors.New("password required")
	}

	if p.Name == "" {
		return errors.New("name required")
	}

	if (p.SSLMode == "disable" || p.SSLMode == "enable") || p.SSLMode == "" {
		return errors.New("sslmode required")
	}

	return nil
}

// RedisConfig holds configuration for Redis connection.
// Used by DB Core service for caching frequently accessed tasks.
//
// Environment Variables:
//   - REDIS_HOST: Redis server host (default: "localhost")
//   - REDIS_PORT: Redis server port (default: "6379")
//   - REDIS_PASSWORD: Redis password (optional, can be empty)
//   - REDIS_DB: Redis database number (default: "10")
type RedisConfig struct {
	RedisHost     string `env:"REDIS_HOST"`     // Redis server host
	RedisPort     string `env:"REDIS_PORT"`     // Redis server port
	RedisPassword string `env:"REDIS_PASSWORD"` // Redis authentication password
	RedisDB       int    `env:"REDIS_DB"`       // Redis database number
}

// validateRedisConfig validates Redis configuration fields.
// Ensures all required fields are present and database number is valid.
//
// Parameters:
//   - r: RedisConfig instance to validate
//
// Returns:
//   - error: Validation error if any field is invalid
func validateRedisConfig(r RedisConfig) error {
	if r.RedisHost == "" {
		return errors.New("redis host required")
	}

	if r.RedisPort == "" {
		return errors.New("redis port required")
	}

	if len([]rune(r.RedisPassword)) == 0 {
		return errors.New("redis password required")
	}

	if r.RedisDB <= 0 {
		return errors.New("redis db required")
	}

	return nil
}

// DBCoreConfig holds configuration for DB Core microservice.
// This service handles database operations and caching for task management.
// Combines PostgreSQL for persistent storage and Redis for caching.
type DBCoreConfig struct {
	DBPostgres PostgresConfig // PostgreSQL database configuration
	DBRedis    RedisConfig    // Redis cache configuration
}

// LoadDBCoreConfig loads DB Core configuration from environment variables.
// Creates and validates both PostgreSQL and Redis configurations.
//
// Returns:
//   - *DBCoreConfig: Loaded configuration with both databases
//   - error: Validation error if any database configuration is invalid
//
// Example:
//
//	cfg, err := LoadDBCoreConfig()
//	if err != nil {
//	    log.Fatal("Failed to load DB config:", err)
//	}
func LoadDBCoreConfig() (*DBCoreConfig, error) {
	postgresConfig := &PostgresConfig{
		Host:     tools.GetEnvOrDefault("POSTGRES_HOST", "localhost"),
		Port:     tools.GetEnvOrDefault("POSTGRES_PORT", "5432"),
		User:     tools.GetEnvOrDefault("POSTGRES_USER", ""),
		Password: tools.GetEnvOrDefault("POSTGRES_PASSWORD", ""),
		Name:     tools.GetEnvOrDefault("POSTGRES_DB", ""),
		SSLMode:  tools.GetEnvOrDefault("POSTGRES_SSLMODE", "disable"),
	}

	if err := validatePostgresConfig(*postgresConfig); err != nil {
		return nil, err
	}

	redisConfig := &RedisConfig{
		RedisHost:     tools.GetEnvOrDefault("REDIS_HOST", "localhost"),
		RedisPort:     tools.GetEnvOrDefault("REDIS_PORT", "6379"),
		RedisPassword: tools.GetEnvOrDefault("REDIS_PASSWORD", ""),
		RedisDB:       tools.ParseIntOrDefault("REDIS_DB", "10"),
	}

	if err := validateRedisConfig(*redisConfig); err != nil {
		return nil, err
	}

	return &DBCoreConfig{
		DBPostgres: *postgresConfig,
		DBRedis:    *redisConfig,
	}, nil
}
