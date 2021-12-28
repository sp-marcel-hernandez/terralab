# terralab

Ranking API on steroids, see [`etc/grpc`](etc/grpc/ranking_service.proto) for service definition.
Research project.

## Requirements

* Go Toolchain +1.17
* `$GOPATH/bin` in your `$PATH` - by default it's `/home/$USER/go/bin` (Linux) or `/Users/$USER/go/bin` (Mac OS)
* GNU Make
* docker
* docker-compose
* JetBrains GoLand

## Research Areas

### Run Go program locally (no docker-compose.yml) with the correct environment

Status: Done

When the backend is not a container in the docker-stack we cannot define its environment variables in the docker-compose.yml file.
However, Makefiles support defining environment variables too:

```makefile
export FOO=bar

# the process WILL receive FOO=bar
run:
	go run github.com/sp-marcel-hernandez/cmd/terralab
```

Better yet, they can be segregated to a separate file and loaded.
This allows easy copy-pasting into GoLand's Run Configuration environment:

```
include .env
export
```


### Deep dive into LocalStack environment variables

Status: Done

Starting from [here](https://github.com/localstack/localstack/blob/master/docker-compose.yml), I cleaned this up until
I had a minimal working LocalStack. See the docker-compose.yml file.

Problematic env vars detected:
```yaml
environment:
  # LocalStack stores useful debug logs here, but only if the env var is defined
  - DATA_DIR=/tmp/localstack/data
  
  # When this env var is not defined its default value is "localstack/localstack:latest" and uses it
  # even if the docker-compose.yml file has a tagged release. My understanding is that this should match
  # the "image:" field.
  - IMAGE_NAME=localstack/localstack:0.13.2
  
  # When this env var is not defined its default value is "localstack_main", and if it doesn't match
  # the container name the internal localstack CLI (at /opt/code/localstack/bin/localstack) doesn't work.
  # This is the automatic name given by docker-compose
  - MAIN_CONTAINER_NAME=terralab_localstack_1
  
  # Makes LocalStack shutdown fast and properly instead of waiting for 20 seconds until Docker kills it by force.
  # Because why do the correct thing by default.
  - SET_TERM_HANDLER=1
```


### Attach Debugging

Status: Done

1. `$ make run`
2. Put breakpoints in source file
3. `Run > Attach to Process (Ctrl+Alt+5) > Select ./out/terralab-dev`
4. Trigger breakpoint (e.g. wait for code to run or make HTTP request)

https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html


### Clean LocalStack provisioning via Terraform without building any Docker images

Status: Pending


### Testing from console with correct environment

Status: Pending


### Testing from GoLand with correct environment and debugging

Status: Pending
