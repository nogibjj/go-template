# go-template
This is a template for building go language applications

## Lesson 3:

Figure out how to build out logic inside of cli (i.e. where do you put it)


## Lesson 2

To run: `go run . --help`
To run: `go run . version


TO DO:

* Get debugger working and learn it
* Write a phrase repeater cli tool
* Build a [cobra tool](https://pkg.go.dev/github.com/spf13/cobra#section-readme).  Link to [official GitHub](https://github.com/spf13/cobra)
* Follow [getting started guide](https://github.com/spf13/cobra-cli/blob/main/README.md): 

`go get -u github.com/spf13/cobra@latest`
`import "github.com/spf13/cobra"`

Install cli generator:
`go install github.com/spf13/cobra-cli@latest`

* Create a new directory:  `mkdir phrase`
* cd into that directory: `cd phrase`
* run go mod init <MODNAME>: `go mod init phrase`

Next init something like this:

```
cd /workspaces/go-template/phrase
cobra-cli init
go run main.go
```

`cobra-cli init --author "Noah Gift"`

## Lesson 1

* Check go version: `go version` and you should get something like this: `go1.19.2 linux/amd64`
* Get gorun installed so you can build go code that runs like a shell script: https://github.com/erning/gorun
* Refer to tutorial:  https://go.dev/doc/tutorial/getting-started
* Create a hello world module:
1. `mkdir hello && cd hello`
2. Enable dependency tracking:  `go mod init example/hello`
3. Paste in example code
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
4. run by doing `go run .` Warning DID NOT WORK! https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code
5.  `mkdir sample-app && cd sample-app`
6.  in a NEW terminal do this:  `go mod init sample-app` (might been my issue before)
7.  paste code in:

```go
package main

import "fmt"

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
}
```