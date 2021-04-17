package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.2.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bs-webserv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bs-webserv version " + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
