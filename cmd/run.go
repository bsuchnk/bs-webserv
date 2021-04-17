package cmd

import (
	"html/template"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var htmlFilePath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the HTTP web server",
	Long: `Start the HTTP web server.
The web server serves HTML file specified by --file flag.`,
	Run: runCmdRun,
}

// runCmdRun starts the HTTP server when the run command is called
func runCmdRun(cmd *cobra.Command, args []string) {
	http.HandleFunc("/", htmlFileHandler)

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	//hide the help flag from Usage
	runCmd.Flags().StringVar(&htmlFilePath, "file", "", "path to the HTML file (required)")
	runCmd.MarkFlagRequired("file")
}

// htmlFileHandler parser the HTML file and executes it
// id the file can't be opened, the function logs an error and exits before execution
func htmlFileHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, r)
}
