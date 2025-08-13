package main

import (
	"fmt"
	"os"
	"todo/internal/service"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var getID string

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"show", "view"},
	Short:   "Get a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if getID == "" {
			fmt.Println("Error: --id is required")
			cmd.Usage()
			return
		}
		taskID, err := uuid.Parse(getID)
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			os.Exit(1)
		}
		tasks, err := service.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			os.Exit(1)
		}
		task, err := tasks.GetTaskById(taskID)
		if err != nil {
			fmt.Println("Error getting task:", err)
			os.Exit(1)
		}
		fmt.Println("Task:", task)
	},
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := service.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			os.Exit(1)
		}
		if len(tasks.GetTasks()) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		tasks.Display()
	},
}

func init() {
	getCmd.Flags().StringVar(&getID, "id", "", "ID of the task to get (required)")
	getCmd.MarkFlagRequired("id")
}
