package cmd

import (
	"fmt"

	"github.com/joelferrier/margarita-mixer/project"
	"github.com/spf13/cobra"
)

func init() {
	profilesCmd.AddCommand(profilesListCmd)
}

var profilesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list availible profiles",
	//Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := project.New() //TODO: handle error
		_ = p.Open()
		fmt.Println("\nprofiles:\n")
		p.PrintProfiles()
	},
}
