package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Hello World")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of aide-family-cli",
	Long:  `All software has versions. This is aide-family-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aide-family-cli v0.0.1 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
