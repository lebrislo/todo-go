package cmd

import (
	"fmt"
	"os"
	"todo-go/controller"

	"github.com/spf13/cobra"
)

var listAllFlag bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks from the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := controller.ListTasks(listAllFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
