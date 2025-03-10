package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"todo_cli/task"
)

// validateDate checks if the due date is in the correct format (YYYY-MM-DD)
func validateDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// validatePriority ensures the priority is between 1 and 5
func validatePriority(priority string) (int, error) {
	p, err := strconv.Atoi(priority)
	if err != nil || p < 1 || p > 5 {
		return 0, fmt.Errorf("priority must be a number between 1 and 5")
	}
	return p, nil
}

// AddTask adds a new task after validating input
func AddTask(description, dueDate, priority string) {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Validate description
	description = strings.TrimSpace(description)
	if description == "" {
		fmt.Println("Error: Task description cannot be empty.")
		return
	}

	// Convert and validate priority
	p, err := validatePriority(priority)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Assign unique ID
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	// Create new task
	newTask := task.Task{
		ID:          newID,
		Description: description,
		Completed:   false,
		Priority:    p, // Now priority is an int
		DueDate:     dueDate,
	}

	// Append task and save
	tasks = append(tasks, newTask)
	err = task.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("✅ Task added successfully!")
}
