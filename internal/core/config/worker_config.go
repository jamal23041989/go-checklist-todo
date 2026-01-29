package config

type WorkerConfig struct {
	KafkaBroker string `env:"KAFKA_BROKER"`
	KafkaTopic  string `env:"KAFKA_TOPIC"`
	LogFilePath string `env:"KAFKA_LOG_FILE_PATH"`
	BatchSize   int    `env:"KAFKA_BATCH_SIZE"`
}

func LoadWorkerConfig() (*WorkerConfig, error) {
	var cfg WorkerConfig

	return &cfg, nil
}
