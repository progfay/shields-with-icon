all: build

build: main.go
	go run ./main.go > README.md 2> Snippets.md

