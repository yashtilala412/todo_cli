package task

import "time"

// Task represents a to-do item
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	DueDate     string    `json:"due_date"`
	Priority    int       `json:"priority"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskManager defines methods for task handling
type TaskManager interface {
	LoadTasks() ([]Task, error)
	SaveTasks([]Task) error
}
