all: build

build: main.go
	go run ./*.go > README.md 2> Snippets.md

