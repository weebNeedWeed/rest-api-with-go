.PHONY: build
build:
	@go build -o ./bin/ecom.exe ./cmd

.PHONY: run
run: build
	@./bin/ecom.exe

.PHONY: test
test:
	@go test ./... -v