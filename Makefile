
.PHONY: build

build: 
	@echo "Building Go with Linux flag"
	@GOOS=linux GOARCH=amd64 go build .
	@echo "Done"