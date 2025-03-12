package cmd

import (
	"fmt"
	"os"
	"strconv"
	"todo-go/csvcontroller"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete <task ID>",
	Short: "Complete a task",
	Long:  `Complete a task from the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Task ID is required")
			return
		}

		taskID, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID")
			return
		}

		err = csvcontroller.CompleteTask(taskID)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
