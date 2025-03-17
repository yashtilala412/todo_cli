package cmd

import (
	"fmt"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// StartReminderSystem prints all due tasks and exits
func StartReminderSystem() {
	printDueTasks()
}

// printDueTasks prints all pending tasks regardless of due date
func printDueTasks() {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	pendingTasksFound := false

	for _, t := range tasks {
		if !t.Completed {
			fmt.Printf("Reminder: Task '%s' is pending! Due Date: %s\n", t.Description, t.DueDate)
			pendingTasksFound = true
		}
	}

	if !pendingTasksFound {
		fmt.Println("No pending tasks.")
	}
}

