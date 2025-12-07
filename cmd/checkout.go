package cmd

import (
	"fmt"
	"mothalali/internal"

	"github.com/spf13/cobra"
)

var createBranch string

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout to a given commitId or branch",
	Long:  `Checkout to a given commitId or branch`,
	Args: func(cmd *cobra.Command, args []string) error {
		if createBranch != "" {
			if len(args) > 0 {
				return fmt.Errorf("accepts 0 arg(s), received %d", len(args))
			}
			return nil
		}
		if len(args) != 1 {
			return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if createBranch != "" {
			fmt.Printf("Creating new branch: %s\n", createBranch)
			// TODO: implement branch creation
		} else {
			commitId := args[0]
			oid := internal.GetOid(commitId)
			fmt.Printf("Checking out to: %s\n", oid)
			internal.Checkout(oid)
		}
	},
}

func init() {
	checkoutCmd.Flags().StringVarP(&createBranch, "branch", "b", "", "Create a new branch")
	rootCmd.AddCommand(checkoutCmd)
}
