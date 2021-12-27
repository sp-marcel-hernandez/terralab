.PHONY: up run down protoc install-tools setup-services dev-build

include .env
export

##########################
### Developer Commands ###
##########################
up: install-tools setup-services protoc run

run: dev-build
	./out/terralab-dev

down:
	docker-compose down -v
	git clean -dfx out etc/terraform

protoc:
	docker-compose run --rm --user $(shell id -u):$(shell id -g) protoc

#######################
### Helper Commands ###
#######################
install-tools:
	go install github.com/google/gops@latest

setup-services:
	docker-compose run --rm starter
	docker-compose run --rm --user $(shell id -u):$(shell id -g) terraform init
	docker-compose run --rm --user $(shell id -u):$(shell id -g) terraform apply -auto-approve

dev-build:
	go build -o out/terralab-dev -gcflags="all=-N -l" github.com/sp-marcel-hernandez/terralab/cmd/terralab
