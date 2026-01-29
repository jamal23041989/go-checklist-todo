// Package tools provides common utility functions used across the application.
// It includes environment variable handling, type parsing, and other technical helpers.
package tools

import (
	"os"
	"strconv"
	"time"
)

// GetEnvOrDefault retrieves environment variable value or returns default value if not set.
// This is useful for configuration with fallback values.
//
// Parameters:
//   - key: Environment variable name
//   - defaultValue: Value to return if environment variable is not set
//
// Returns:
//   - Environment variable value if set, otherwise defaultValue
//
// Example:
//
//	port := GetEnvOrDefault("HTTP_PORT", "8080")
func GetEnvOrDefault(key, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}
	return defaultValue
}

// ParseDurationOrDefault parses environment variable as time.Duration or returns default.
// Supports duration formats like "30s", "5m", "1h".
// Returns 30 seconds as safe default if parsing fails.
//
// Parameters:
//   - key: Environment variable name containing duration string
//   - defaultValue: Default duration string if environment variable is not set
//
// Returns:
//   - Parsed duration or safe default (30s) if parsing fails
//
// Example:
//
//	timeout := ParseDurationOrDefault("HTTP_TIMEOUT", "30s")
func ParseDurationOrDefault(key, defaultValue string) time.Duration {
	envValue := os.Getenv(key)
	if envValue == "" {
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

// ParseIntOrDefault parses environment variable as integer or returns default.
// Returns 0 as safe default if parsing fails.
//
// Parameters:
//   - key: Environment variable name containing integer value
//   - defaultValue: Default integer value as string if environment variable is not set
//
// Returns:
//   - Parsed integer or 0 if parsing fails
//
// Example:
//
//	batchSize := ParseIntOrDefault("BATCH_SIZE", "100")
func ParseIntOrDefault(key, defaultValue string) int {
	envValue := os.Getenv(key)
	if envValue == "" {
		if parsedValue, err := strconv.Atoi(defaultValue); err == nil {
			return parsedValue
		}
		return 0
	}

	if parsedValue, err := strconv.Atoi(envValue); err == nil {
		return parsedValue
	}
	return 0
}
