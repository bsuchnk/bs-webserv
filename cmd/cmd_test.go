package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testCases = []struct {
	testName    string
	fileName    string
	fileContent string
}{
	{
		testName:    "one_line",
		fileName:    "testHtmlFile1",
		fileContent: "Hello, World!",
	},
	{
		testName: "multi_line",
		fileName: "testHtmlFile2",
		fileContent: `Multi
Line
File`,
	},
	{
		testName: "working_html",
		fileName: "testHtmlFile2",
		fileContent: `<head><title>Title</title></head>
<body>
	<div class="page-heading">
        <h1>Page Heading</h1>
	</div>
</body>`,
	},
}

// TestCorrectHtmlFileHandler ests if http server responds with
// Status OK and the page's body is equal to the HTML file's content
// if the HTML file provided by the user is correct
func TestCorrectHtmlFileHandler(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
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

			// simulate providing path to the HTML file by the user
			htmlFilePath = f.Name()

			// create GET request for the http server
			request, err := http.NewRequest("GET", "localhost:8080", nil)
			if err != nil {
				t.Fatal(err)
			}
			responseRecorder := httptest.NewRecorder()

			// send the request to the http server
			htmlFileHandler(responseRecorder, request)

			// get the result
			res := responseRecorder.Result()
			defer res.Body.Close()

			// fail if the server responded with incorrect status
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, got %v", res.Status)
			}

			// fail if page's body is incorrect
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != testCase.fileContent {
				t.Errorf("HTML file content incorrect\nexpected:\n\"%v\"\ngot:\n\"%v\"", string(body), testCase.fileContent)
			}
		})
	}
}

// TestNonexistentHtmlFileHandler tests if http server responds with
// Status Not Found, when file path provided by the user is incorrect
func TestNonexistentHtmlFileHandler(t *testing.T) {
	testCase := testCases[0]

	f, err := os.CreateTemp("", testCase.fileName)
	if err != nil {
		t.Fatal(err)
	}
	// remove the file right afterwards
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		t.Fatal(err)
	}

	// simulate providing path to the HTML file by the user
	htmlFilePath = f.Name()

	// create GET request for the http server
	request, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()

	// send the request to the server
	htmlFileHandler(responseRecorder, request)

	// get the result
	res := responseRecorder.Result()
	defer res.Body.Close()

	// fail if the server didn't respond with Status Not Found
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Response Status incorrect\nexpected:\n\"%v\"\ngot:\n\"%v\"", http.StatusNotFound, res.Status)
	}
}
