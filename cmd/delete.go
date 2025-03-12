package cmd

import (
	"fmt"
	"os"
	"strconv"
	"todo-go/csvcontroller"

	"github.com/spf13/cobra"
)

var deleteAllFlag bool

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID] | --all",
	Short: "delete task(s)",
	Long:  `Delete a specific task by providing its ID or delete all tasks using the --all flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteAllFlag {
			err := csvcontroller.DeleteAll()

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else if len(args) > 0 {
			taskID, err := strconv.Atoi(args[0])

			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid task ID")
				return
			}

			err = csvcontroller.DeleteTask(taskID)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {

		}
	},
}
