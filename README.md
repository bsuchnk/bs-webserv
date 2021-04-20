# bs-webserv

bs-webserv is a simple CLI program for starting an HTTP web server serving an HTML file.

## Table of Contents
- [Installing](#installing)
- [Running](#running)
- [Examples](#examples)
- [Tests](#tests)
- [Tools used](#tools-used)

## Installing
    $ go get github.com/bsuchnk/bs-webserv
    go: downloading github.com/bsuchnk/bs-webserv v1.0.2
    
## Running
    $ bs-webserv help
    bs-webserv is a simple CLI program which starts an http web server.
    The server serves an HTML file selected by the user.

    Usage:
      bs-webserv [command]

    Available Commands:
      help        Help about any command
      run         Start the HTTP web server
      version     Print the version of bs-webserv

    Use "bs-webserv [command] --help" for more information about a command.
    
## Examples
### - Existing HTML file
    $ bs-webserv run --file index.html
    2021/04/19 21:27:58 Listening on :8080...

![example_html](https://user-images.githubusercontent.com/75221970/115292979-07cc7a80-a157-11eb-90aa-37fb942aa254.png)

### - Nonexistent HTML file
    $ bs-webserv run --file nonexistent.html
    2021/04/19 21:34:56 Listening on :8080...

![nonexistent_example](https://user-images.githubusercontent.com/75221970/115293135-3f3b2700-a157-11eb-91ab-dc3457c2bd59.png)

## Tests
    $ cd $GOPATH/pkg/mod/github.com/bsuchnk/bs-webserv@v1.0.2
    $ go test -v ./cmd
    === RUN   TestCorrectHtmlFileHandler
    === RUN   TestCorrectHtmlFileHandler/one_line
    === RUN   TestCorrectHtmlFileHandler/multi_line
    === RUN   TestCorrectHtmlFileHandler/working_html
    --- PASS: TestCorrectHtmlFileHandler (0.02s)
        --- PASS: TestCorrectHtmlFileHandler/one_line (0.01s)
        --- PASS: TestCorrectHtmlFileHandler/multi_line (0.01s)
        --- PASS: TestCorrectHtmlFileHandler/working_html (0.01s)
    === RUN   TestNonexistentHtmlFileHandler
    --- PASS: TestNonexistentHtmlFileHandler (0.00s)
    PASS
    ok      github.com/bsuchnk/bs-webserv/cmd       0.100s

## Tools used:
- **Go**
- **Cobra**
