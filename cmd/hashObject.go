/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha1"
	"fmt"
	"mothalali/internal"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "creating SHA-1 of a file that given",
	Long: `read the given file and saves it in the object store
	Example:
      mothalali hash-object <filename>`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		obj_type := args[1]

		fmt.Println("given args:", path, obj_type)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("File not found:", path)
			os.Exit(1)
		} else {
			b, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}
			var header []byte
			header = fmt.Appendf(header, "%s %d\x00", obj_type, len(b))
			b = append(header, b...)
			hash := sha1.New()
			hash.Write(b)
			file_hash := hash.Sum(nil)
			dir := filepath.Join(internal.ObjectsDir, fmt.Sprintf("%x", file_hash[:1]))
			file_path := filepath.Join(dir, fmt.Sprintf("%x", file_hash[1:]))

			// Ensure the directory exists
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("Error creating directory:", err)
				os.Exit(1)
			}
			err = os.WriteFile(file_path, b, 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				os.Exit(1)
			}
			fmt.Printf("file_path: %s\n", file_path)
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
