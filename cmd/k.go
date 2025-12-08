/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mothalali/internal"

	"github.com/spf13/cobra"
)

// kCmd represents the k command
var kCmd = &cobra.Command{
	Use:   "k",
	Short: "The visualization tool will draw all refs and all the commits pointed by the refs",
	Long:  `The visualization tool will draw all refs and all the commits pointed by the refs`,
	Run: func(cmd *cobra.Command, args []string) {
		s := make(map[string]struct{})
		for ref := range internal.IterRefs() {
			s[ref.OID] = struct{}{}
		}
		internal.IterCommitsAndParents(s)
	},
}

func init() {
	rootCmd.AddCommand(kCmd)
}
