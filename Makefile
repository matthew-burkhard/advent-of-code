build:
	go build

run:
	make build
	./main

tidy:
	go mod tidy