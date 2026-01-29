package domains

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jamal23041989/go-checklist-todo/internal/core/apperr"
)

// Event represents a system event for Kafka logging.
// It captures user actions with timestamps for audit and monitoring.
//
// Business rules:
// - Action must be one of: create, delete, update, list, done
// - ID is auto-generated UUID for event tracking
// - Timestamp is set automatically on creation
// - TaskID is optional (not required for list actions)
type Event struct {
	ID        uuid.UUID  `json:"id"`
	TaskID    *uuid.UUID `json:"task_id,omitempty"`
	Action    string     `json:"action"`
	Timestamp time.Time  `json:"timestamp"`
}

// NewCreateEvent creates an event for task creation
func NewCreateEvent(taskID uuid.UUID) *Event {
	return &Event{
		ID:        uuid.New(),
		TaskID:    &taskID,
		Action:    "create",
		Timestamp: time.Now(),
	}
}

// NewDeleteEvent creates an event for task deletion
func NewDeleteEvent(taskID uuid.UUID) *Event {
	return &Event{
		ID:        uuid.New(),
		TaskID:    &taskID,
		Action:    "delete",
		Timestamp: time.Now(),
	}
}

// NewUpdateEvent creates an event for task update
func NewUpdateEvent(taskID uuid.UUID) *Event {
	return &Event{
		ID:        uuid.New(),
		TaskID:    &taskID,
		Action:    "update",
		Timestamp: time.Now(),
	}
}

// NewDoneEvent creates an event for task completion
func NewDoneEvent(taskID uuid.UUID) *Event {
	return &Event{
		ID:        uuid.New(),
		TaskID:    &taskID,
		Action:    "done",
		Timestamp: time.Now(),
	}
}

// NewListEvent creates an event for task list retrieval
func NewListEvent() *Event {
	return &Event{
		ID:        uuid.New(),
		TaskID:    nil,
		Action:    "list",
		Timestamp: time.Now(),
	}
}

// Validate validates event according to business rules
func (e *Event) Validate() error {
	validActions := map[string]bool{
		"create": true,
		"delete": true,
		"update": true,
		"list":   true,
		"done":   true,
	}

	if !validActions[e.Action] {
		return errors.New(apperr.ErrInvalidAction)
	}

	if e.ID == uuid.Nil {
		return errors.New(apperr.ErrIDRequired)
	}

	return nil
}
