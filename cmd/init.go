/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mothalali/internal"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize mothalali git in your directory",
	Long: `initialize mothalali git in your directory.
Example:
  mothalali init`,
	Run: func(cmd *cobra.Command, args []string) {
		absPath, err := filepath.Abs(internal.GitDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting absolute path: %v\n", err)
			os.Exit(1)
		}

		if _, err := os.Stat(internal.GitDir); os.IsNotExist(err) {
			if err := os.Mkdir(internal.GitDir, 0700); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Initialized empty Mothalali repository in %s\n", absPath)
		} else {
			fmt.Printf("Mothalali repository already exists in %s\n", absPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
