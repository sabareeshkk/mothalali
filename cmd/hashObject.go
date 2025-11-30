/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mothalali/internal"

	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "creating SHA-1 of a file that given",
	Long: `read the given file and saves it in the object store
	Example:
      mothalali hash-object <filename>
      mothalali hash-object <filename> -e blob`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		obj_type := args[1]

		fmt.Println("given args:", path, obj_type)
		if file_path, err := internal.HashObject(path, obj_type); err != nil {
			cmd.PrintErrf("Error: %v", err)
		} else {
			fmt.Println("successfully created:", file_path)
		}
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashObjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashObjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
