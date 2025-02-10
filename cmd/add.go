package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "Add a new task %s\n", args[0])

	},
}
