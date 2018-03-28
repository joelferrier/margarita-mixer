package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "perform operations on profiles",
	//Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loadProject()
	},
}
