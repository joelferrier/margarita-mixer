package cmd

import (
	"fmt"

	"github.com/joelferrier/margarita-mixer/project"
	"github.com/spf13/cobra"
)

func init() {
	profileCmd.AddCommand(profileListCmd)
}

var profileListCmd = &cobra.Command{
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
