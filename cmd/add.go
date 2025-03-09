package cmd

import (
	"todo-go/csvcontroller"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		csvcontroller.AddTask(args[0])
	},
}
