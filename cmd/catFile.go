package cmd

import (
	"fmt"
	"mothalali/internal"

	"github.com/spf13/cobra"
)

var expected string

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "shows the content of the file with the provided SHA-1 hash",
	Long: `Shows the content of the file with the provided SHA-1 hash
	Example:
		mothalali cat-file <SHA-1_hash/tag-name>
		mothalali cat-file <SHA-1_hash/tag-name> -e blob
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Expected value:", expected)
		hashOrName := args[0]
		oid := internal.GetOid(hashOrName)
		content, err := internal.ReadObject(oid, expected)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(content)) // print anyways
	},
}

func init() {
	catFileCmd.Flags().StringVarP(&expected, "expected", "e", "", "Expected blob type")
	rootCmd.AddCommand(catFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
