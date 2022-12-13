# go-template
This is a template for building go language applications


## Lesson 4: Continue the basic tutorial

https://go.dev/doc/tutorial/handle-errors

## Lesson 3:

## Expanded Hello World with Modules

* you'll create two modules.

* `mkdir -p greetings && cd greetings`
* `go mod init example.com/greetings`
* `touch greetings.go` and paste in:
* btw, you CANNOT DO THIS NO MAIN: `go run .`

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

make new module

* `cd ..`
* `mkdir -p hello && cd hello`
* `go mod init example.com/hello`
* `touch hello.go` and paste this in:

```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

* `go mod edit -replace example.com/greetings=../greetings`
* `go mod tidy`
* `go run .`




Interactive playground is amazing!:  https://go.dev/play/

### Build Calc

1. `mkdir -p calc && cd calc`
2. `go mod init example/calc`
3. `touch calc.go`

```go
// Build a calculator

package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println("Sum: ", a+b)
}
```

Will build out:

```bash
@noahgift ➜ /workspaces/go-template/calc (main ✗) $ go run .
Enter two numbers: 
2 2
Sum:  4
```







rehash tutorial!

1. `mkdir hello && cd hello`
2. `go mod init example/hello`
3.  `touch hello.go`
4.  paste this in:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

5. `go run .`

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