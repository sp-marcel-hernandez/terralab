# terralab

## Requirements

* Go Toolchain +1.17
* GNU Make
* docker
* docker-compose
* JetBrains GoLand


## Attach Debugging

1. `$ make run`
2. Put breakpoints in source file
3. `Run > Attach to Process (Ctrl+Alt+5) > Select ./out/terralab`
4. Trigger breakpoint (e.g. wait for code to run or make HTTP request)

https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html


## Open Fronts

* [X] Attach-debugging to local running process
* [ ] LocalStack provisioning via Terraform
* [ ] Testing from console with correct environment
* [ ] Testing from GoLand with correct environment and debugging
