## ----------------------------------------------------------------------
## htmx-server
##
## following you can find all available tooling steps to
## install dependencies, run development mode, build the project,
## format the project or run the build
## ----------------------------------------------------------------------

.PHONY: help
help:		## Show this help
	@awk '/(^##)|^[a-z].*##/' Makefile

.PHONY: install
install:	## Install dependencies
	go mod tidy

.PHONY: build
build:		## Build the project
	go build .

.PHONY: generate
generate:	## generates templates from templ files
	templ generate

.PHONY: format
format:		## Format the project
	go fmt .

.PHONY: run
run:		## runs the build server
	go run .

.PHONY: serve
serve:		## build server and start it running
serve: build run
