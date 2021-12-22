.PHONY: run clean install-tools build stack-up stack-down

include .env
export

##########################
### Developer Commands ###
##########################
run: install-tools stack-up build
	./out/terralab

clean: stack-down
	rm -rf out

###################
### CI Commands ###
###################

#######################
### Helper Commands ###
#######################
install-tools:
	go install github.com/google/gops@latest

build:
	go build -o out/terralab -gcflags="all=-N -l" github.com/sp-marcel-hernandez/terralab/cmd/terralab

stack-up:
	docker-compose run --rm starter

stack-down:
	docker-compose down -v
