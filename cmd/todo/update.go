package main

import (
	"fmt"
	"os"
	"todo/internal/service"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	updateID          string
	updateTitle       string
	updateDescription string
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"edit", "modify"},
	Short:   "Update an existing task",
	Run: func(cmd *cobra.Command, args []string) {
		if updateID == "" {
			fmt.Println("Error: --id is required")
			cmd.Usage()
			return
		}
		if updateTitle == "" {
			fmt.Println("Error: --t (title) is required")
			cmd.Usage()
			return
		}
		taskID, err := uuid.Parse(updateID)
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			os.Exit(1)
		}
		tasks, err := service.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			os.Exit(1)
		}
		exist, err := tasks.GetTaskById(taskID)
		if err != nil {
			fmt.Println("Error getting task:", err)
			os.Exit(1)
		}
		desc := updateDescription
		if desc == "" {
			desc = exist.Description
		}
		tasks.UpdateTask(taskID, updateTitle, desc)
		err = service.SaveTasks(tasks.GetTasks())
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Println("Task updated successfully.")
	},
}

func init() {
	updateCmd.Flags().StringVar(&updateID, "id", "", "ID of the task to update (required)")
	updateCmd.Flags().StringVarP(&updateTitle, "title", "t", "", "New title of the task (required)")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "New description of the task")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("title")
}
