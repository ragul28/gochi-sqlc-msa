.PHONY: build run mod create_migrate migrate_up migrate_down

# Go toolchain
run:
	go run cmd/usersrv/main.go

build_musl:
	mkdir ./build/bin/ || true
	CGO_ENABLED=0 go build -v -ldflags="-s -w" -o=./build/bin/. ./cmd/...

build_glibc:
	mkdir ./build/bin/ || true
	go build -v -ldflags="-s -w" -o=./build/bin/. ./cmd/...

mod:
	go mod tidy
	go mod verify

# Database & migration
DB_URL="postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"

db_access:
	docker exec -it postgres psql $(DB_URL)

sqlc:
	sqlc generate

create_migrate:
	migrate create -ext=sql -dir=internal/db/migrations -seq init

migrate_up:
	migrate -database $(DB_URL) -path internal/db/migrations -verbose up 

migrate_down:
	migrate -database $(DB_URL) -path internal/db/migrations -verbose  down