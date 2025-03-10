package cmd

import (
	"todo-go/csvcontroller"

	"github.com/spf13/cobra"
)

var allFlag bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks from the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		csvcontroller.ListTasks(allFlag)
	},
}
