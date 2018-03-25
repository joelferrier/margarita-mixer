package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(profilesCmd)
}

var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "perform operations on profiles",
	//Long:  ``,
}
