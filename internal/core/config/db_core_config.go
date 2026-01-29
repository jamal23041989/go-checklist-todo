package config

import (
	"errors"

	"github.com/jamal23041989/go-checklist-todo/internal/core/tools"
)

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Name     string `env:"POSTGRES_DB"`
	SSLMode  string `env:"POSTGRES_SSLMODE"`
}

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

type RedisConfig struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPort     string `env:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`
}

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

type DBCoreConfig struct {
	DBPostgres PostgresConfig
	DBRedis    RedisConfig
}

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
