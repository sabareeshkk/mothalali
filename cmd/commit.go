/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mothalali/internal"

	"github.com/spf13/cobra"
)

var message string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit your changes",
	Long: `command that will accept a commit message,
	snapshot the current directory using mothalali write-tree and save the resulting object`,
	Example: "mothalai commit -m 'Initial commit'",
	Run: func(cmd *cobra.Command, args []string) {
		oid, err := internal.Commit(message)
		if err != nil {
			panic(err)
		}
		fmt.Println("Commit successful:", oid)
	},
}

func init() {
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
	if err := commitCmd.MarkFlagRequired("message"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(commitCmd)
}
