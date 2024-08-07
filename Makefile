# Makefile

# Variables
GOOS := linux
GOARCH := amd64
CGO_ENABLED := 0
OUTPUT := bootstrap
ZIPFILE := myLambda.zip
STATIC_DIR := static

# Default target
all:
	@echo "Building..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) go build -tags lambda.norpc -o $(OUTPUT) main.go
	@echo "Creating zip file..."
	zip -r $(ZIPFILE) $(OUTPUT)
	@echo "Adding static directory to zip..."
	zip -r $(ZIPFILE) $(STATIC_DIR)
	@echo "Done!"

.PHONY: all