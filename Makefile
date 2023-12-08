build:
	go mod download
	go vet -v
	go test -v
	CGO_ENABLED=0 sudo go build -o /usr/local/bin/pm




