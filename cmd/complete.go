package cmd

import (
	"fmt"
	"strconv"

	"todo_cli/task"
)

// CompleteTask marks a task as completed by ID.
func CompleteTask(idStr string) {
	// Convert ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID. Please enter a valid number.")
		return
	}

	// Load tasks from cache (no need to read file again)
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Find and update the task
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	// Save updated tasks back to file
	if err := task.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task ID %d marked as completed!\n", id)
}
