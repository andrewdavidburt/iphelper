GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=main
ZIP_NAME=$(BINARY_NAME).zip

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(ZIP_NAME)

aws-lambda-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v; \
	zip -j $(ZIP_NAME) $(BINARY_NAME)