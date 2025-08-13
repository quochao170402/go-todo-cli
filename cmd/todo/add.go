package main

import (
	"fmt"
	"os"
	"todo/internal/service"

	"github.com/spf13/cobra"
)

var (
	addTitle       string
	addDescription string
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"new", "create"},
	Short:   "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if addTitle == "" {
			fmt.Println("Error: --t (title) is required")
			cmd.Usage()
			return
		}
		tasks, err := service.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			os.Exit(1)
		}
		tasks.AddTask(addTitle, addDescription)
		err = service.SaveTasks(tasks.GetTasks())
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	},
}

func init() {
	addCmd.Flags().StringVarP(&addTitle, "title", "t", "", "Title of the task (required)")
	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "Description of the task")
	addCmd.MarkFlagRequired("t")
}
