build:
	go build -o bin/catfacts

run: build
	./bin/catfacts
