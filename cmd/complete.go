package cmd

import (
	"fmt"
	"strconv"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
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
	fmt.Printf("\n%-4s %-10s %-30s %-10s %-12s\n", "ID", "Status", "Task", "Priority", "Due Date")

	// Print Completed Task Details
	t := tasks[id-1]
	fmt.Printf("%-4d %-10s %-30s %-10d %-12s\n", id, "Completed", t.Description, t.Priority, t.DueDate)
}

