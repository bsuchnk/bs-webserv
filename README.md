# bs-webserv

## Installing
    go get github.com/bsuchnk/bs-webserv
    
## Running
    $ bs-webserv help
    bs-webserv is a simple CLI program which starts an http web server.
    The server serves an HTML file selected by the user.

    Usage:
      bs-webserv [command]

    Available Commands:
      help        Help about any command
      run         Start the HTTP web server
      version     Print the version number of bs-webserv

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
