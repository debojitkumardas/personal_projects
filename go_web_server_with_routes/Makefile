PROJECT_NAME=go_web_server_with_routes

SRC_DIR=cmd/$(PROJECT_NAME)
INT_DIR=internal
BIN_DIR=bin

GO=go

SRC=$(SRC_DIR)/main.go
BIN=$(BIN_DIR)/$(PROJECT_NAME)

.PHONY: all build clean

all: build

build:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(SRC)

run: build
	echo "Message: execute curl_localhost.sh in a seperate terminal"
	@./$(BIN)

clean:
	@rm -rf $(BIN)
