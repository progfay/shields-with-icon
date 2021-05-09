all: clean generator run

clean:
	rm README.md Snippets.md

generator: main.go color.go icon.go
	go build -o=generator ./*.go

run: main.go
	./generator > README.md 2> Snippets.md
