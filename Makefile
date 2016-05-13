GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD) -v -o bin/todo-view

install:
	$(GOINSTALL)

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/todo-view
	rm -rf bin/todo-view

test:
	$(GOTEST) -v -cover

.PHONY: all clean
