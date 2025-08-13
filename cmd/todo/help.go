package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help",
	Short:   "Show help message",
	Aliases: []string{"--help", "-h"},
	Run: func(cmd *cobra.Command, args []string) {
		printUserGuide()
	},
}

func printUserGuide() {
	fmt.Println("Usage: todo [command] [options]")
	fmt.Println("Commands:")
	for _, c := range rootCmd.Commands() {
		aliasesStr := ""
		if len(c.Aliases) > 0 {
			aliasesStr = fmt.Sprintf(" (aliases: %s)", strings.Join(c.Aliases, ", "))
		}
		fmt.Printf("  %-10s%s - %s\n", c.Use, aliasesStr, c.Short)
	}
}
