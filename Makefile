build:
	@go build -o bin/portfolio-cms cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/portfolio-cms

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations -seq $(name)
	
migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down