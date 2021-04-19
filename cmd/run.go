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

// runCmdRun starts the http server
// when the run command is called
func runCmdRun(cmd *cobra.Command, args []string) {
	http.HandleFunc("/", htmlFileHandler)

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// htmlFileHandler parses the HTML file and executes it.
// if the file can't be opened, the function sends appropriate status and response
func htmlFileHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlFilePath)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<h1>404 - File Not Found</h1>"))
		return
	}

	t.Execute(w, r)
}

func init() {
	rootCmd.AddCommand(runCmd)

	//hide the help flag from Usage
	runCmd.Flags().StringVar(&htmlFilePath, "file", "", "path to the HTML file (required)")
	runCmd.MarkFlagRequired("file")
}
