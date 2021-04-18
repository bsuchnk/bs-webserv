package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// func TestVersionOutput(t *testing.T) {
// 	buffer := bytes.NewBufferString("")
// 	versionCmd.SetOut(buffer)

// 	runCmd.Execute()

// 	out, err := ioutil.ReadAll(buffer)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := "bs-webserv v"
// 	if string(out) != expected {
// 		t.Fatalf("expected \"%s\", got \"%s\"", expected, string(out))
// 	}
// }

var testCases = []struct {
	fileName    string
	fileContent string
}{
	{
		"testHtmlFile1",
		"Hello, World!",
	},
	{
		"testHtmlFile2",
		`Multi
Line
File`,
	},
	{
		"testHtmlFile2",
		`<head><title>Title</title></head>
<body>
</body>`,
	},
}

func TestHtmlFileHandler(t *testing.T) {
	for _, testCase := range testCases {
		// create temporary HTML file
		f, err := os.CreateTemp("", testCase.fileName)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(f.Name()) // clean up

		// write content to the HTML file and close it
		if _, err := f.Write([]byte(testCase.fileContent)); err != nil {
			t.Fatal(err)
		}
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}

		// try to open the file through the http server
		htmlFilePath = f.Name()

		request, err := http.NewRequest("GET", "localhost:8080", nil)
		if err != nil {
			t.Fatal(err)
		}
		responseRecorder := httptest.NewRecorder()

		htmlFileHandler(responseRecorder, request)

		res := responseRecorder.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK, got %v", res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) == testCase.fileContent {
			t.Errorf("HTML file content incorrect\nexpected:\n\"%v\"\ngot:\n\"%v\"", string(body), testCase.fileContent)
		}
	}
}
