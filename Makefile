GOCMD=go
FIELDALIGNMENT_CMD=fieldalignment
GOBUILD=$(GOCMD) build
BINARY_NAME=kubectl-envsecret
MAIN_FILE=main.go

vet:
	@echo "Vet the code"
	$(GOCMD) vet ./...

fieldalignment-check:
	@echo "Running fieldalignment"
	$(FIELDALIGNMENT_CMD) ./...

fieldalignment:
	@echo "Running fieldalignment fix to optimize struct sizes"
	$(FIELDALIGNMENT_CMD) -fix ./...

test:
	@echo "Running tests"
	$(GOCMD) test ./...

format:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

build: vet
	@echo "Building..."
	$(GOCMD) build -o $(BINARY_NAME) $(MAIN_FILE)

fix: format fieldalignment

build-prod: test fieldalignment-check build

run: build
	./$(BINARY_NAME) $(CMD) $(ARGS) $(FLAGS)


.PHONY: build build-prod run format vet fieldalignment-check fieldalignment test fix
