GOCMD=go
GOOS_UBUNTU=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_PATH=bin
BINARY_NAME=ezgo

all: hello test build
hello:
	echo "This is Makefile of ezgo"
build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -v
test:
	$(GOTEST) -v ./core/...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)/$(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -v
	./$(BINARY_PATH)/$(BINARY_NAME)
# Cross compilation
build-linux:
	$(GOOS_UBUNTU) $(GOBUILD) -o $(BINARY_UNIX) -v