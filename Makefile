BUILD_DIR=bin
BINARY=gateway

all: clean build run

build:deps
	go build -o $(BUILD_DIR)/$(BINARY) ./cmd/api/main.go

clean:
	rm -rf $(BUILD_DIR)

deps:
	go mod tidy

run:
	./$(BUILD_DIR)/$(BINARY)

re: clean build run

.PHONY: build clean deps run re
