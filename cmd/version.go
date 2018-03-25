package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of margarita-mixer",
	Long:  `All software has versions. This is margarita-mixer's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("margarita-mixer v0.1.0")
	},
}
