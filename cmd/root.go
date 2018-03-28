package cmd

import (
	"fmt"
	"os"

	"github.com/joelferrier/margarita-mixer/project"
	"github.com/spf13/cobra"
)

var proj project.Project

var rootCmd = &cobra.Command{
	Use:   "margarita-mixer",
	Short: "margarita-mixer is a LiME module build tool.",
	Long: `margarita-mixer is LiME module build tool which
uses docker to compile LiME kernel modules and creates
a searchable repository of kernel modules.`,
	Hidden: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadProject() {
	proj, err := project.New()
	if err != nil {
		panic(err)
	}
	err = proj.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println("project opened")

}
