package config

import (
	"errors"

	"github.com/jamal23041989/go-checklist-todo/internal/core/tools"
)

type WorkerConfig struct {
	KafkaBroker string `env:"KAFKA_BROKER"`
	KafkaTopic  string `env:"KAFKA_TOPIC"`
	LogFilePath string `env:"KAFKA_LOG_FILE_PATH"`
	BatchSize   int    `env:"KAFKA_BATCH_SIZE"`
}

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
