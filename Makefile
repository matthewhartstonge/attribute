# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODEP=dep ensure -v
BINARY_NAME=attribute
BINARY_BUILD_PATH=./cmd/attribute/main.go
BINARY_DIST_PATH=./dist/
BINARY_DARWIN=$(BINARY_NAME)_darwin
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME).exe
BUILD_FLAGS=-ldflags="-s -w"

all: build-all
build:
		$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIST_PATH)$(BINARY_NAME) -v $(BINARY_BUILD_PATH)
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GODEP)

build-all: build-darwin	build-linux	build-windows

# Cross compilation
build-darwin:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIST_PATH)$(BINARY_DARWIN) -v $(BINARY_BUILD_PATH)
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIST_PATH)$(BINARY_UNIX) -v $(BINARY_BUILD_PATH)
build-windows:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIST_PATH)$(BINARY_WINDOWS) -v $(BINARY_BUILD_PATH)
