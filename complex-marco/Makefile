init:
	go mod init github.com/noahgift/complex-marco

install:
	go get github.com/urfave/cli/v2

clean:
	rm -rf marco-polo

format:
	go fmt ./...

lint:
	golangci-lint run

build:
	go build -o complex-marco

run:
	./marco-polo "Marco"

all: clean install format lint build run