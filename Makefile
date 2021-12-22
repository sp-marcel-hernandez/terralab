.PHONY: build install-tools run

export SP_CONVERT=UUooOoOoooUUUOoOoOOO

build:
	go build -o out/terralab github.com/sp-marcel-hernandez/terralab/cmd/terralab

install-tools:
	go install github.com/go-delve/delve/cmd/dlv@latest

run: install-tools
	go run github.com/sp-marcel-hernandez/terralab/cmd/terralab
