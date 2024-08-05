.PHONY: run build

run:
	go run cmd/main/main.go

build:
	go build -o bin/addigy cmd/main/main.go

build-move:
	go build -o /usr/local/bin/addigy cmd/main/main.go && chmod +x /usr/local/bin/addigy
