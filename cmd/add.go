package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// validatePriority ensures the priority is a valid integer >= 1
func validatePriority(priority string) (int, error) {
	p, err := strconv.Atoi(priority)
	if err != nil || p < 1 {
		return 0, fmt.Errorf("priority must be a positive number")
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

	// Validate due date using Go's time.Parse (YYYY-MM-DD format)
	_, err = time.Parse("2006-01-02", dueDate)
	if err != nil {
		fmt.Println("Error: Invalid due date format. Please use YYYY-MM-DD.")
		return
	}

	// Convert and validate priority
	p, err := validatePriority(priority)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Assign unique ID
	var newID int
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID
	}
	newID++ // Increment to get the next ID

	// Create new task
	newTask := task.Task{
		ID:          newID,
		Description: description,
		Completed:   false,
		Priority:    p, // Priority is an integer
		DueDate:     dueDate,
	}

	// Append task and save
	tasks = append(tasks, newTask)
	err = task.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task added successfully!")
}
