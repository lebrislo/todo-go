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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "TODO-GO !\n")
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "u", "LE BRIS Loris", "author name for copyright attribution")

	rootCmd.AddCommand(addCmd)

	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "List all tasks")
	rootCmd.AddCommand(listCmd)

	rootCmd.AddCommand(completeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
