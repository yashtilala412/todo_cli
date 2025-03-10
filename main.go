package main

import (
	"fmt"
	"os"

	"git.pride.improwised.dev/Onboarding-2025/Yash-Tilala/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add|list]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 { // Ensure enough arguments
			fmt.Println("Usage: go run main.go add <description> <due_date> <priority>")
			return
		}
		description := os.Args[2]
		dueDate := os.Args[3]
		priority := os.Args[4]

		cmd.AddTask(description, dueDate, priority)

	case "list":
		cmd.ListTasks()

	case "complete":
		cmd.CompleteTask(os.Args[2])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <task_id>")
			return
		}
		cmd.DeleteTask(os.Args[2])

	default:
		fmt.Println("Invalid command. Available commands: add, list")
	}
}

