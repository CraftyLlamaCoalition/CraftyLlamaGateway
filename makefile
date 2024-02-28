BINARY=build/gateway

build:
	go build -o $(BINARY) ./cmd/api/main.go

clean:
	rmm -rf $(BINARY)

deps:
	go mod tidy

run:
	./$(BINARY)

re: clean build run

.PHONY: build clean deps run re
