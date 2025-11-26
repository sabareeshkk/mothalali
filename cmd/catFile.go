/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"mothalali/internal"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var expected string

func parseGitObject(object []byte) (objType string, size int, content []byte, err error) {
	spaceIdx := bytes.IndexByte(object, ' ')
	nullIdx := bytes.IndexByte(object, 0)
	if spaceIdx == -1 || nullIdx == -1 || nullIdx < spaceIdx {
		err = fmt.Errorf("invalid git object format")
		return
	}

	objType = string(object[:spaceIdx])
	sizeStr := string(object[spaceIdx+1 : nullIdx])
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return
	}
	content = object[nullIdx+1:]
	return
}

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "shows the content of the file with the provided SHA-1 hash",
	Long: `Shows the content of the file with the provided SHA-1 hash
	Example:
		mothalali cat-file <SHA-1_hash>
		mothalali cat-file <SHA-1_hash> -e blob
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Expected value:", expected)
		sha1_hash := args[0]
		dir := filepath.Join(internal.ObjectsDir, sha1_hash[:2])
		filepath := filepath.Join(dir, sha1_hash[2:])
		content, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		objType, size, content, err := parseGitObject(content)
		fmt.Println("objType:", objType, "size:", size)
		if err != nil {
			fmt.Println("Error parsing git object:", err)
			return
		}
		if expected != "" && objType != expected {
			fmt.Println("Unexpected object type:", objType)
			return
		}
		if objType == "blob" {
			fmt.Println(string(content))
		} else {
			fmt.Printf("Object type: %s, size: %d bytes\n", objType, size)
		}
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
