acceptance:
	go test -v ./cmd/...

acceptance-test-http:
	go test -v ./cmd/httpserver/...

unit-tests:
	go test -v ./specifications/...
lint:
	golangci-lint run ./...

run-http-server:
	go run cmd/httpserver/*.go

run: run-http-server