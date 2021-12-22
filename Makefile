.PHONY: install-tools build run stack-up stack-down

export SP_CONVERT=UUooOoOoooUUUOoOoOOO


install-tools:
	go install github.com/go-delve/delve/cmd/dlv@latest

build:
	go build -o out/terralab github.com/sp-marcel-hernandez/terralab/cmd/terralab

run: install-tools
	go run github.com/sp-marcel-hernandez/terralab/cmd/terralab

stack-up:
	docker-compose run --rm starter

stack-down:
	docker-compose down -v
