package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  `Complete a task from the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "Complete a task\n")
	},
}
