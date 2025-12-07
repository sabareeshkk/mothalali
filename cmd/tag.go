package cmd

import (
	"mothalali/internal"

	"github.com/spf13/cobra"
)

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "create a tag",
	Long: `create a tag	 on a commitId 
	This is to name a commit for easier reference.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		tagName := args[0]
		if len(args) == 2 {
			commitId := args[1]
			oid := internal.GetOid(commitId)
			internal.CreateTag(tagName, oid)
		} else {
			internal.CreateTag(tagName, "")
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
