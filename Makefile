.PHONY: build run mod

# Go toolchain
run:
	go run cmd/usersrv/main.go

build_glibc:
	mkdir ./build/bin/ || true
	go build -v -ldflags="-s -w" -o=./build/bin/. ./cmd/...

mod:
	go mod tidy
	go mod verify
