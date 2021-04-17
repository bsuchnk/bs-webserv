package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bs-webserv",
	Short: "bs-webserv is a simple CLI program which starts a HTTP web server.",
	Long: `bs-webserv is a simple CLI program which starts a HTTP web server.
The server serves selected HTML file.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Hide --help flag from Usage
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.PersistentFlags().Lookup("help").Hidden = true
}
