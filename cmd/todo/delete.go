package main

import (
	"fmt"
	"os"
	"todo/internal/service"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var deleteID string

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "remove", "del"},
	Short:   "Delete a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID == "" {
			fmt.Println("Error: --id is required")
			cmd.Usage()
			return
		}
		taskID, err := uuid.Parse(deleteID)
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			os.Exit(1)
		}
		tasks, err := service.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			os.Exit(1)
		}
		tasks.RemoveTask(taskID)
		err = service.SaveTasks(tasks.GetTasks())
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted successfully.")
	},
}

func init() {
	deleteCmd.Flags().StringVar(&deleteID, "id", "", "ID of the task to delete (required)")
	deleteCmd.MarkFlagRequired("id")
}
