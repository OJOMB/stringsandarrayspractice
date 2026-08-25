.PHONY: test run vet

test:
	go test ./internal/...

run:
	go run .

vet:
	go vet ./...

lint:
	golangci-lint run