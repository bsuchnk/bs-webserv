package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of bs-webserv",
	Run:   versionCmdRun,
}

// versionCmdRun prints out the program's version
// when the version command is called
func versionCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("bs-webserv version " + version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
