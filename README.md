# go-template
This is a template for building go language applications

## Lesson 2

TO DO:

* Get debugger working and learn it
* Write a phrase repeater cli tool

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