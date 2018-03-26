package cmd

import (
	"fmt"

	"github.com/joelferrier/margarita-mixer/builder"
	"github.com/joelferrier/margarita-mixer/builder/docker"
	"github.com/joelferrier/margarita-mixer/project"
	"github.com/spf13/cobra"
)

func init() {
	buildCmd.AddCommand(buildAllCmd)
}

var buildAllCmd = &cobra.Command{
	Use:   "all",
	Short: "build all configured profiles",
	//Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := project.New() //TODO: handle error
		_ = p.Open()
		fmt.Println("project opened")

		var b builder.Builder
		b = docker.NewBuilder()

		b.Configure(p.Builders["docker"])
		_ = b.Setup()
		_ = b.Build()
		_ = b.Extract()
		_ = b.Cleanup()
	},
}
