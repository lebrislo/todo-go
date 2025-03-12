package cmd

import (
	"fmt"
	"os"
	"todo-go/csvcontroller"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a new task",
	Long:  `Add a new task to the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Task description is required")
			return
		}

		err := csvcontroller.AddTask(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
