/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mothalali/internal"
	"os"

	"github.com/spf13/cobra"
)

// writeTreeCmd represents the writeTree command
var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "writing the entire directory tree",
	Long: `write-tree is a command that writes the entire directory tree to the database
	Example:
		mothalali write-tree <directory_path>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("writeTree called")
		directoryPath := args[0]
		fmt.Printf("Writing directory tree from %s\n", directoryPath)
		tree, err := internal.WriteTree(directoryPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Directory tree written to %s\n", tree)
	},
}

func init() {
	rootCmd.AddCommand(writeTreeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeTreeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeTreeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
