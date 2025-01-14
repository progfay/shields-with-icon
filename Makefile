all: clean generator run

clean:
	rm -f README.md Snippets.md

generator: *.go
	go build -o=generator ./*.go

run: main.go
	./generator
