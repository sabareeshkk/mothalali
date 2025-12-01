/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mothalali/internal"

	"github.com/spf13/cobra"
)

// readTreeCmd represents the readTree command
var readTreeCmd = &cobra.Command{
	Use:   "read-tree",
	Short: "restore to the commit-id",
	Long: `This command restores the working directory to the state of the specified commit-id.
	Example:
		mothalali read-tree <commit-id>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readTree called")
		commitID := args[0]
		internal.ReadTree(commitID)
	},
}

func init() {
	rootCmd.AddCommand(readTreeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readTreeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readTreeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
