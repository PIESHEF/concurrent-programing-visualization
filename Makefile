BINARY_NAME=hello-world

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} src/main.go

run:
	./bin/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -r bin