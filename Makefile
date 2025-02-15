BINARY_NAME=$(HOME)/go/bin/pm

VERSION=v1.0.2
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# Flags para o linker
LDFLAGS=-ldflags "-X 'main.version=$(VERSION)' -X 'main.commit=$(COMMIT)' -X 'main.date=$(DATE)'"

build:
	@echo "Compilando $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME)

clean:
	@echo "Limpando..."
	rm -f $(BINARY_NAME)

all: clean build