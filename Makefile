BUILD_FOLDER  ?= build
BINARY_NAME   ?= merkletree-implementation
CONFIG_PATH   ?= config.yaml

.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(BUILD_FOLDER)/$(BINARY_NAME) 

.PHONY: test
test:
	@echo "Running tests..."
	@go clean -testcache && go test ./... -cover

.PHONY: run
run: 
	@echo "Running the application..."
	@go run main.go
.PHONY: clean	
clean:
	@echo "Cleaning up..."
	@go clean
	@rm -rf $(BUILD_FOLDER)

