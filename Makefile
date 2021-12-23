.PHONY: up down install-tools setup-services protoc debug-build

include .env
export

##########################
### Developer Commands ###
##########################
up: install-tools setup-services protoc debug-build
	./out/terralab

down:
	docker-compose down -v
	rm -rf \
	etc/terraform/.terraform \
	etc/terraform/.terraform.lock.hcl \
	etc/terraform/terraform.tfstate \
	etc/terraform/terraform.tfstate.backup \
	out

#######################
### Helper Commands ###
#######################
install-tools:
	go install github.com/google/gops@latest

setup-services:
	docker-compose run --rm starter
	docker-compose run --rm --user $(shell id -u):$(shell id -g) terraform init
	docker-compose run --rm --user $(shell id -u):$(shell id -g) terraform apply -auto-approve

protoc:
	docker-compose run --rm --user $(shell id -u):$(shell id -g) protoc

debug-build:
	go build -o out/terralab -gcflags="all=-N -l" github.com/sp-marcel-hernandez/terralab/cmd/terralab
