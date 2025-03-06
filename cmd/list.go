package cmd

import (
	"fmt"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// ListTasks prints all tasks with their status.
func ListTasks() {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Your Tasks:")
	for _, t := range tasks {
		status := "❌ Pending"
		if t.Completed {
			status = "✅ Completed"
		}
		fmt.Printf("[%s] %s - Priority: %s - Due: %s\n", status, t.Description, t.Priority, t.DueDate)
	}
}
