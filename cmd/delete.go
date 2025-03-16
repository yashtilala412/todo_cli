package cmd

import (
	"fmt"
	"strconv"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/task"
)

// DeleteTask removes a task by its ID
func DeleteTask(taskID string) {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("Invalid task ID:", err)
		return
	}

	found := false
	var updatedTasks []task.Task

	// Remove the task with the given ID
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue // Skip the task to be deleted
		}
		updatedTasks = append(updatedTasks, t)
	}

	if !found {
		fmt.Println("Task not found!")
		return
	}

	err = task.SaveTasks(updatedTasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task deleted successfully!")
}
