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
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: build
build:		## Build the project
	go build -o dist/ .

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
serve: generate run

.PHONY: watch
watch:		## watch go and templ files, rebuild and restart server. Based of "fd" and "entr" unix utility tools
	fd --exclude '*_templ.go' | entr -rs 'make generate && sh watch.sh main.go'
