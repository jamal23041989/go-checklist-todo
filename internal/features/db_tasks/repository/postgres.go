// Package repository provides data access layer for task management.
// It implements the TaskRepository interface using PostgreSQL as the storage backend.
// This package contains SQL queries and database operations for tasks.
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jamal23041989/go-checklist-todo/internal/core/domains"
	"github.com/jamal23041989/go-checklist-todo/internal/infrastructure/postgres"
)

// TaskRepository defines the contract for task data operations.
// This interface follows the Repository pattern and can be implemented
// with different storage backends (PostgreSQL, MongoDB, etc.).
//
// All methods accept context for cancellation and timeout handling.
type TaskRepository interface {
	// Create inserts a new task into the database.
	// Returns the created task with generated fields populated.
	Create(ctx context.Context, task domains.Task) (*domains.Task, error)

	// Get retrieves a task by its unique identifier.
	// Returns nil if task is not found.
	Get(ctx context.Context, id uuid.UUID) (*domains.Task, error)

	// List retrieves all tasks from the database.
	// Returns empty slice if no tasks exist.
	List(ctx context.Context) ([]domains.Task, error)

	// Update modifies an existing task by its ID.
	// Returns the updated task with new values.
	Update(ctx context.Context, id uuid.UUID) (*domains.Task, error)

	// Delete removes a task from the database by its ID.
	// Returns the deleted task for confirmation.
	Delete(ctx context.Context, id uuid.UUID) (*domains.Task, error)
}

// taskRepository implements TaskRepository interface using PostgreSQL.
// It contains SQL queries and database operations for task management.
// This struct is private to enforce dependency injection through interface.
type taskRepository struct {
	db *postgres.DB // Database connection pool
}

// NewTaskRepository creates a new PostgreSQL task repository.
// It accepts a database connection and returns a TaskRepository interface.
//
// Parameters:
//   - db: PostgreSQL database connection pool
//
// Returns:
//   - TaskRepository: Repository interface for task operations
//
// Example:
//
//	db, err := postgres.NewDB(cfg)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	repo := NewTaskRepository(db)
//	task, err := repo.Create(ctx, task)
func NewTaskRepository(db *postgres.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(ctx context.Context, task domains.Task) (*domains.Task, error) {
	return nil, nil
}

func (r *taskRepository) Get(ctx context.Context, id uuid.UUID) (*domains.Task, error) {
	return nil, nil
}

func (r *taskRepository) List(ctx context.Context) ([]domains.Task, error) {
	return nil, nil
}

func (r *taskRepository) Update(ctx context.Context, id uuid.UUID) (*domains.Task, error) {
	return nil, nil
}

func (r *taskRepository) Delete(ctx context.Context, id uuid.UUID) (*domains.Task, error) {
	return nil, nil
}
