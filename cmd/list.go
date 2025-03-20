package cmd

import (
	"fmt"
	"sort"
	"time"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// ListTasks prints all tasks with their status.
// ListTasks prints all tasks with their status.
func ListTasks() {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	sort.SliceStable(tasks, func(primary, secondary int) bool {

		if tasks[primary].Priority != tasks[secondary].Priority {
			return tasks[primary].Priority < tasks[secondary].Priority
		}

		dueDatePrimary, errPrimary := time.Parse("2006-01-02", tasks[primary].DueDate)
		dueDateSecondary, errSecondary := time.Parse("2006-01-02", tasks[secondary].DueDate)

		// If due date format is invalid, keep original order
		if errPrimary != nil || errSecondary != nil {
			return false
		}

		return dueDatePrimary.Before(dueDateSecondary)
	})

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	// Print Header Row
	fmt.Printf("\n%-4s %-10s %-30s %-10s %-12s\n", "SrNo", "Status", "Task", "Priority", "Due Date")

	// Print Task List
	for i, t := range tasks {
		status := "Pending"
		if t.Completed {
			status = "Completed"
		}
		fmt.Printf("%-4d %-10s %-30s %-10d %-12s\n", i+1, status, t.Description, t.Priority, t.DueDate)
	}
}
