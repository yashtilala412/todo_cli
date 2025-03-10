package cmd

import (
	"fmt"
	"sort"
	"time"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// ListTasks prints all tasks with their status.
func ListTasks() {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	sort.SliceStable(tasks, func(i, j int) bool {
		// First, compare priority (higher is better)
		if tasks[i].Priority != tasks[j].Priority {
			return tasks[i].Priority < tasks[j].Priority
		}

		// If priority is the same, compare due dates (earlier is better)
		dateI, errI := time.Parse("2006-01-02", tasks[i].DueDate)
		dateJ, errJ := time.Parse("2006-01-02", tasks[j].DueDate)

		// If due date format is invalid, keep original order
		if errI != nil || errJ != nil {
			return false
		}

		return dateI.Before(dateJ) // Earlier due date first
	})

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Your Tasks:")
	for i, t := range tasks {
		status := "Pending"
		if t.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %-12s %-30s Priority: %d  Due: %s\n", i+1, status, t.Description, t.Priority, t.DueDate)
	}

}

