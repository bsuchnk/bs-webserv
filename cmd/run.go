package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var htmlFilePath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the HTTP web server",
	Long: `Start the HTTP web server.
The web server serves HTML file specified by --file flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVar(&htmlFilePath, "file", "", "path to the HTML file (required)")
	runCmd.MarkFlagRequired("file")
}
