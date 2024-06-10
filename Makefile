GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=kubectl-envsecret
MAIN_FILE=main.go

test:
	@echo "Running tests"
	$(GOCMD) test ./...

format:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

build:
	@echo "Building..."
	$(GOCMD) build -o $(BINARY_NAME) $(MAIN_FILE)

build-prod: test format build

run: build
	./$(BINARY_NAME) $(CMD) $(ARGS) $(FLAGS)


.PHONY: build build-prod run format
