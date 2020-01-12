# fe-message-server
records messages sent from frontend clients to persistance

# Setup local evironment
## Prerequisite
- Download Go Tool via https://golang.org/ and install them base on your operation system.
  
## IDE
- I prefer to use Vitual Studio Code. So Download them and install https://code.visualstudio.com/
- Next, install plugin which is call Go
- Next, CMD + Ship + P for mac to seach `Go: Install/Update` and it will prompt the tools out of the box. So select all of them to install and for this step will take a while.

## GOPATH
- check GOPATH: go to command line and typing go env GOPATH
- Next, export PATH:$PATH:$(go env GOPATH)/bin
- Next, (optional) export GOPATH=$(go env GOPATH)
- more details here https://golang.org/doc/code.html

## Create workspace or Go initializer module
- create workspace directory and then create go module initializer by cli and typing to `go mod init github.developer.allianz.io/hexalite/fe-messsaging-server`

# Try with Go sandbox
- https://play.golang.org/


go test -v -cover ./...
go build -v -ldflags"-X 'main.Version=v1.0.0' -X 'main.Time=$(date)' -X 'main.User=$(id -u -n)'"

{
  "app": {
    "server": "127.0.0.1",
    "port": "8080",
    "pprofPort": "8081"
  },
  "mongo": {
    "server": "mongo:27017",
    "database": "messaging_db",
    "collection": [
      "client_message"
    ]
  },
  "elastic": {
    "server": "http://10.17.175.106:12101/client_log_index/_bulk"
  }
}