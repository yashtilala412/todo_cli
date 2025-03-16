package cmd

import (
	"fmt"
	"time"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

func StartReminderSystem() {

	printDueTasks()
	go func() {
		for {
			time.Sleep(24 * time.Hour) // Wait for a day
			printDueTasks()            // Check and display tasks again
		}
	}()
}
func printDueTasks() {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	today := time.Now().Format("2006-01-02")
	dueTasksFound := false

	for _, t := range tasks {
		if !t.Completed && t.DueDate == today {
			fmt.Printf("Reminder: Task '%s' is due today! (%s)\n", t.Description, t.DueDate)
			dueTasksFound = true
		}
	}
	if !dueTasksFound {
		fmt.Println("No tasks are due today.")
	}
}
