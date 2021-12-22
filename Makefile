.PHONY: run install-tools build stack-up stack-down

export SP_CONVERT=UUooOoOoooUUUOoOoOOO

run: install-tools stack-up build
	./out/terralab

install-tools:
	go install github.com/google/gops@latest

build:
	go build -o out/terralab -gcflags="all=-N -l" github.com/sp-marcel-hernandez/terralab/cmd/terralab

clean: stack-down
	rm -rf out

stack-up:
	docker-compose run --rm starter

stack-down:
	docker-compose down -v
