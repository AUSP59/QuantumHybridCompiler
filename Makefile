build:
	go build -o bin/qhc ./cmd/qhc

lint:
	golangci-lint run

test:
	go test ./...

serve:
	go run ./internal/server
