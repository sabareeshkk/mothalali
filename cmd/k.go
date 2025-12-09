/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mothalali/internal"

	"github.com/spf13/cobra"
)

// kCmd represents the k command
var kCmd = &cobra.Command{
	Use:   "k",
	Short: "The visualization tool will draw all refs and all the commits pointed by the refs",
	Long:  `The visualization tool will draw all refs and all the commits pointed by the refs`,
	Run: func(cmd *cobra.Command, args []string) {
		dot := "digraph commits {\n"
		s := make(map[string]struct{})
		for ref := range internal.IterRefs() {
			dot += fmt.Sprintf("\"%s\" [shape=note]\n", ref.Name)
			dot += fmt.Sprintf("\"%s\" -> \"%s\"\n", ref.Name, ref.OID)
			s[ref.OID] = struct{}{}
		}
		internal.IterCommitsAndParents(s, &dot)
		dot += "}"
		fmt.Println(dot)
		internal.ShowDot(dot)
	},
}

func init() {
	rootCmd.AddCommand(kCmd)
}
