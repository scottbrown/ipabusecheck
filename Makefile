.DEFAULT_GOAL := build

app.name := ipabusecheck
app.repo := github.com/scottbrown/$(app.name)

build.dir := .build
dist.dir  := .dist

.PHONY: build
build:
	go build -o $(build.dir)/$(app.name) $(app.repo)/cmd

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(build.dir) $(dist.dir)

.PHONY: release
release: get-version
	GOOS=linux GOARCH=amd64 go build -o $(build.dir)/linux-amd64/dumpcft $(app.repo)/cmd
	GOOS=linux GOARCH=arm64 go build -o $(build.dir)/linux-arm64/dumpcft $(app.repo)/cmd
	GOOS=darwin GOARCH=amd64 go build -o $(build.dir)/darwin-amd64/dumpcft $(app.repo)/cmd
	GOOS=darwin GOARCH=arm64 go build -o $(build.dir)/darwin-arm64/dumpcft $(app.repo)/cmd
	GOOS=windows GOARCH=amd64 go build -o $(build.dir)/windows-amd64/dumpcft $(app.repo)/cmd
	GOOS=windows GOARCH=arm64 go build -o $(build.dir)/windows-arm64/dumpcft $(app.repo)/cmd
	mkdir -p $(dist.dir)
	tar cfz $(dist.dir)/$(app.name)_$(VERSION)_linux_amd64.tar.gz -C $(build.dir)/linux-amd64 .
	tar cfz $(dist.dir)/$(app.name)_$(VERSION)_linux_arm64.tar.gz -C $(build.dir)/linux-arm64 .
	tar cfz $(dist.dir)/$(app.name)_$(VERSION)_darwin_amd64.tar.gz -C $(build.dir)/darwin-amd64 .
	tar cfz $(dist.dir)/$(app.name)_$(VERSION)_darwin_arm64.tar.gz -C $(build.dir)/darwin-arm64 .
	tar cfz $(dist.dir)/$(app.name)_$(VERSION)_windows_amd64.tar.gz -C $(build.dir)/windows-amd64 .
	tar cfz $(dist.dir)/$(app.name)_$(VERSION)_windows_arm64.tar.gz -C $(build.dir)/windows-arm64 .

.PHONY: get-version
get-version:
ifndef VERSION
	@echo "Provide a VERSION to continue."; exit
endif
