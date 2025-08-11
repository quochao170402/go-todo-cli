package main

import (
	"fmt"

	"todo/data"
	"todo/service"

	"os"

	"github.com/google/uuid"
)

func main() {
	if len(os.Args) < 2 {
		printUserGuide()
		return
	}

	command := normalizeCommand(os.Args[1])

	tasks, err := service.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch command {
	case "add":
		handleAdd(tasks)
	case "update":
		handleUpdate(tasks)
	case "get":
		handleGet(tasks)
	case "delete":
		handleDelete(tasks)
	case "list":
		handleList(tasks)
	case "complete":
		handleComplete(tasks)
	case "help":
		printUserGuide()
	default:
		fmt.Println("Unknown command:", command)
	}
}

func normalizeCommand(cmd string) string {
	switch cmd {
	case "ls":
		return "list"
	case "rm":
		return "delete"
	case "--a":
		return "add"
	case "--u":
		return "update"
	default:
		return cmd
	}
}

func handleAdd(tasks data.TodoList) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: add <title> <description>")
		return
	}
	title := os.Args[2]
	description := ""
	if len(os.Args) > 3 {
		description = os.Args[3]
	}
	tasks.AddTask(title, description)
	err := service.SaveTasks(tasks.GetTasks())
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task added successfully.")
}

func handleUpdate(tasks data.TodoList) {
	if len(os.Args) < 5 {
		fmt.Println("Usage: update <id> <title> <description>")
		return
	}
	id := os.Args[2]
	title := os.Args[3]
	description := ""
	if len(os.Args) > 4 {
		description = os.Args[4]
	}
	taskID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid task ID:", err)
		return
	}
	exist, err := tasks.GetTaskById(taskID)
	if err != nil {
		fmt.Println("Error getting task:", err)
		return
	}
	if len(description) == 0 {
		description = exist.Description
	}
	tasks.UpdateTask(taskID, title, description)
	err = service.SaveTasks(tasks.GetTasks())
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task updated successfully.")
}

func handleGet(tasks data.TodoList) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: get <id>")
		return
	}
	id := os.Args[2]
	taskID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid task ID:", err)
		return
	}
	task, err := tasks.GetTaskById(taskID)
	if err != nil {
		fmt.Println("Error getting task:", err)
		return
	}
	fmt.Println("Task:", task)
}

func handleDelete(tasks data.TodoList) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: delete <id>")
		return
	}
	id := os.Args[2]
	taskID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid task ID:", err)
		return
	}
	tasks.RemoveTask(taskID)
	err = service.SaveTasks(tasks.GetTasks())
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task deleted successfully.")
}

func handleList(tasks data.TodoList) {
	if len(tasks.GetTasks()) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	tasks.Display()
}

func handleComplete(tasks data.TodoList) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: complete <id>")
		return
	}
	id := os.Args[2]
	taskID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid task ID:", err)
		return
	}
	tasks.MarkTaskDone(taskID)
	err = service.SaveTasks(tasks.GetTasks())
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task marked as done successfully.")
}

func printUserGuide() {
	fmt.Println("Usage: [command] [options]")
	fmt.Println("Commands:")
	fmt.Println("  add <title> <description> - Add a new task")
	fmt.Println("  update <id> <title> <description> - Update an existing task")
	fmt.Println("  get <id> - Get a task by ID")
	fmt.Println("  delete <id> - Delete a task by ID")
	fmt.Println("  list - List all tasks")
	fmt.Println("  complete <id> - Mark a task as done")
	fmt.Println("  help - Show this help message")
}
