package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-go",
	Short: "TODO-GO is a lightweight task manager",
	Long:  `TODO-GO is a lightweight task manager built with Go and Cobra.`,
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "u", "LE BRIS Loris", "author name for copyright attribution")
	/* add */
	rootCmd.AddCommand(addCmd)
	/* list */
	listCmd.Flags().BoolVarP(&listAllFlag, "all", "a", false, "List all tasks")
	rootCmd.AddCommand(listCmd)
	/* complete */
	rootCmd.AddCommand(completeCmd)
	/* delete */
	deleteCmd.Flags().BoolVarP(&deleteAllFlag, "all", "a", false, "Delete all tasks")
	rootCmd.AddCommand(deleteCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
