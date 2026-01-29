// Package config provides configuration management for all microservices.
// Each service has its own configuration structure with environment variable loading
// and validation. Configurations are loaded at application startup.
package config

import (
	"errors"

	"github.com/jamal23041989/go-checklist-todo/internal/core/tools"
)

// WorkerConfig holds configuration for Kafka Worker microservice.
// This service consumes events from Kafka and logs them to files.
//
// Environment Variables:
//   - KAFKA_BROKER: Kafka broker address (default: "localhost:9092")
//   - KAFKA_TOPIC: Kafka topic to consume events from (default: "task_events")
//   - KAFKA_LOG_FILE_PATH: Path to log file for event storage (default: "./logs/worker.log")
//   - KAFKA_BATCH_SIZE: Number of events to process in batch (default: "100")
type WorkerConfig struct {
	KafkaBroker string `env:"KAFKA_BROKER"`        // Kafka broker address
	KafkaTopic  string `env:"KAFKA_TOPIC"`         // Kafka topic name
	LogFilePath string `env:"KAFKA_LOG_FILE_PATH"` // Log file path
	BatchSize   int    `env:"KAFKA_BATCH_SIZE"`    // Batch processing size
}

// LoadWorkerConfig loads Worker configuration from environment variables.
// Sets up Kafka connection and logging configuration with default values.
//
// Returns:
//   - *WorkerConfig: Loaded configuration for Kafka worker
//   - error: Validation error if required fields are missing or invalid
//
// Example:
//
//	cfg, err := LoadWorkerConfig()
//	if err != nil {
//	    log.Fatal("Failed to load worker config:", err)
//	}
func LoadWorkerConfig() (*WorkerConfig, error) {
	workerConfig := &WorkerConfig{
		KafkaBroker: tools.GetEnvOrDefault("KAFKA_BROKER", "localhost:9092"),
		KafkaTopic:  tools.GetEnvOrDefault("KAFKA_TOPIC", "kafka"),
		LogFilePath: tools.GetEnvOrDefault("KAFKA_LOG_FILE_PATH", "/var/log/kafka-client.log"),
		BatchSize:   tools.ParseIntOrDefault("KAFKA_BATCH_SIZE", "10"),
	}

	if err := workerConfig.validateWorkerConfig(); err != nil {
		return nil, err
	}

	return workerConfig, nil
}

// validateWorkerConfig validates Worker configuration fields.
// Ensures all required Kafka settings are present and valid.
//
// Parameters:
//   - w: WorkerConfig instance to validate
//
// Returns:
//   - error: Validation error if any field is invalid
func (w WorkerConfig) validateWorkerConfig() error {
	if w.KafkaBroker == "" {
		return errors.New("kafka broker required")
	}

	if w.KafkaTopic == "" {
		return errors.New("kafka topic required")
	}

	if w.LogFilePath == "" {
		return errors.New("kafka log file path required")
	}

	return nil
}
