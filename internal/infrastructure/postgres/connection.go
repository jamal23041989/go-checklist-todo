// Package postgres provides PostgreSQL connection management for the application.
// It handles connection pooling, configuration, and health checks for PostgreSQL database.
// This package focuses only on technical database connectivity without business logic.
package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jamal23041989/go-checklist-todo/internal/core/config"
)

const (
	MaxConns          = 10               // Maximum number of connections in pool
	MinConns          = 2                // Minimum number of connections in pool
	HealthCheckPeriod = 30 * time.Second // Health check interval for connections
)

// DB represents a PostgreSQL database connection pool.
// It provides high-performance connection management using pgxpool.
//
// Usage:
//
//	db, err := NewDB(cfg)
//	if err != nil {
//	    log.Fatal("Failed to connect to database:", err)
//	}
//	defer db.Close()
//
//	// Use in repository
//	rows, err := db.Pool.Query(ctx, "SELECT * FROM tasks")
type DB struct {
	Pool *pgxpool.Pool // Connection pool for PostgreSQL
}

// NewDB creates a new PostgreSQL database connection pool with the given configuration.
// It parses the connection string, configures pool settings, and validates connectivity.
//
// Parameters:
//   - cfg: Database configuration containing host, port, credentials, and SSL settings
//
// Returns:
//   - *DB: Configured database connection pool
//   - error: Connection or configuration error
//
// Example:
//
//	cfg := config.DBConfig{...}
//	db, err := NewDB(cfg)
//	if err != nil {
//	    return fmt.Errorf("database setup failed: %w", err)
//	}
func NewDB(cfg config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBPostgres.Host,
		cfg.DBPostgres.Port,
		cfg.DBPostgres.User,
		cfg.DBPostgres.Password,
		cfg.DBPostgres.Name,
		cfg.DBPostgres.SSLMode,
	)

	ctx := context.Background()

	// Используем pgxpool.Config для детальной настройки
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}

	// Настройки connection pool
	poolConfig.MaxConns = MaxConns
	poolConfig.MinConns = MinConns
	poolConfig.HealthCheckPeriod = HealthCheckPeriod

	// Создаем connection pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &DB{Pool: pool}, nil
}

// Close gracefully closes the database connection pool.
// Should be called during application shutdown to release resources.
//
// Example:
//
//	defer db.Close()
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
