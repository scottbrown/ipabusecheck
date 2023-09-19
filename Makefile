.DEFAULT_GOAL := build

app.name := ipabusecheck
app.repo := github.com/scottbrown/$(app.name)

build.dir := .build

.PHONY: build
build:
	go build -o $(build.dir)/$(app.name) $(app.repo)/cmd

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...
