package domains

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jamal23041989/go-checklist-todo/internal/core/apperr"
)

// Task represents a task in the checklist system.
// It encapsulates all business rules and validation logic for task management.
//
// Business rules:
// - Name is required and max 200 characters
// - Description is optional and max 1000 characters
// - ID must be a valid UUID
// - CreatedAt is set automatically on creation
// - UpdatedAt is updated on any modification
type Task struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	IsDone      bool       `json:"is_done" db:"is_done"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

// NewTask creates a new task with validation.
// Returns error if name is empty or too long, or description is too long.
func NewTask(name, description string) (*Task, error) {
	if err := validateTaskFields(name, description); err != nil {
		return nil, err
	}

	return &Task{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}, nil
}

// Validate validates all task fields according to business rules.
// Ensures ID is valid UUID and name/description meet constraints.
func (t *Task) Validate() error {
	if t.ID == uuid.Nil {
		return errors.New(apperr.ErrIDRequired)
	}

	if err := validateTaskFields(t.Name, t.Description); err != nil {
		return err
	}

	return nil
}

// MarkDone marks the task as completed and updates UpdatedAt timestamp.
func (t *Task) MarkDone() error {
	t.IsDone = true
	now := time.Now()
	t.UpdatedAt = &now

	return nil
}

// MarkNotDone marks the task as not completed and updates UpdatedAt timestamp.
func (t *Task) MarkNotDone() error {
	t.IsDone = false
	now := time.Now()
	t.UpdatedAt = &now

	return nil
}

// Update updates task name and description with validation.
// Updates UpdatedAt timestamp on successful modification.
func (t *Task) Update(name, description string) error {
	if err := validateTaskFields(name, description); err != nil {
		return err
	}

	t.Name = name
	t.Description = description
	now := time.Now()
	t.UpdatedAt = &now

	return nil
}

// validateTaskFields validates name and description according to business rules.
// Private helper function to avoid code duplication.
func validateTaskFields(name, description string) error {
	if name == "" {
		return errors.New(apperr.ErrNameRequired)
	}

	if len([]rune(name)) > 200 {
		return errors.New(apperr.ErrNameTooLong)
	}

	if len([]rune(description)) > 1000 {
		return errors.New(apperr.ErrDescriptionTooLong)
	}

	return nil
}
