package cmd

import (
	"mothalali/internal"

	"github.com/spf13/cobra"
)

var commitID string
var ref string

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Long: `Show the commit logs for the current repository.
You can specify a starting commit ID using the --commitid flag.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if ref != "" {
			commitID = internal.GetOid(ref)
		}
		internal.GetCommit(commitID)
	},
}

func init() {
	logCmd.Flags().StringVarP(&commitID, "commitid", "c", "", "Commit object ID")
	logCmd.Flags().StringVarP(&ref, "ref", "r", "", "Ref name")
	rootCmd.AddCommand(logCmd)
}
