package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bs-webserv",
	Short: "bs-webserv is a simple CLI program which starts an HTTP web server.",
	Long: `bs-webserv is a simple CLI program which starts an HTTP web server.
The server serves one HTML file.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

}
