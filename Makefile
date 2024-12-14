.PHONY: build
build:
	@go build -o ./bin/ecom.exe ./cmd

.PHONY: run
run: build
	@./bin/ecom.exe

.PHONY: test
test:
	@go test ./... -v

MIGRATIONS_FOLDER=./cmd/migrate/migrations
MIGRATE_ENTRYPOINT=./cmd/migrate/main.go

.PHONY: add-migration
add-migration:
	@migrate create -ext sql -dir $(MIGRATIONS_FOLDER) -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@go run $(MIGRATE_ENTRYPOINT) up

.PHONY: migrate-down
migrate-down:
	@go run $(MIGRATE_ENTRYPOINT) down