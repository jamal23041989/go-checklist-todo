package tools

import (
	"os"
	"strconv"
	"time"
)

func GetEnvOrDefault(key, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}
	return defaultValue
}

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
