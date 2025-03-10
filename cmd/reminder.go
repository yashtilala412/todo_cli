package cmd

import (
	"fmt"
	"time"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// StartReminderSystem runs a background goroutine to check for due tasks daily
func StartReminderSystem() {
	go func() {
		for {
			tasks, err := task.LoadTasks()
			if err != nil {
				fmt.Println("Error loading tasks:", err)
				return
			}

			today := time.Now().Format("2006-01-02") // Get today's date in YYYY-MM-DD format

			for _, t := range tasks {
				// Compare due date with today's date
				if !t.Completed && t.DueDate == today {
					fmt.Printf("Reminder: Task '%s' is due today! (%s)\n", t.Description, t.DueDate)
				}
			}

			// Sleep for 24 hours before checking again
			time.Sleep(24 * time.Hour)
		}
	}()
}
