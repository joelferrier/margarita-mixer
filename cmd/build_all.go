package cmd

import (
	"github.com/joelferrier/margarita-mixer/builder"
	"github.com/joelferrier/margarita-mixer/builder/docker"
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
		var b builder.Builder
		b, err := docker.NewBuilder()
		if err != nil {
			panic(err)
		}

		b.Configure(proj.Builders["docker"])
		_ = b.Setup(proj.Profiles["amzn1"])
		_ = b.Build()
		_ = b.Extract()
		_ = b.Cleanup()
	},
}
