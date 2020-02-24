appname := logcli

sources := $(wildcard *.go logcli/*.go)
cli = github.com/florentsolt/log/logcli

build = GOOS=$(1) GOARCH=$(2) CGO_ENABLED=0 go build -o build/$(appname) $(cli)
tar = cd build && tar -cvzf $(1)_$(2).tar.gz $(appname) && rm $(appname)
zip = cd build && zip $(1)_$(2).zip $(appname) && rm $(appname)

.PHONY: all windows darwin linux clean

all: windows darwin linux

dev: $(sources)
	go build -o build/$(appname) $(cli)

lint:
	$(HOME)/go/bin/golint ./...

clean:
	rm -rf build/

##### LINUX BUILDS #####
linux: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: $(sources)
	$(call build,linux,386,)
	$(call tar,linux,386)

build/linux_amd64.tar.gz: $(sources)
	$(call build,linux,amd64,)
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: $(sources)
	$(call build,linux,arm,)
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: $(sources)
	$(call build,linux,arm64,)
	$(call tar,linux,arm64)

##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_amd64.tar.gz

build/darwin_amd64.tar.gz: $(sources)
	$(call build,darwin,amd64,)
	$(call tar,darwin,amd64)

##### WINDOWS BUILDS #####
windows: build/windows_386.zip build/windows_amd64.zip

build/windows_386.zip: $(sources)
	$(call build,windows,386,.exe)
	$(call zip,windows,386,.exe)

build/windows_amd64.zip: $(sources)
	$(call build,windows,amd64,.exe)
	$(call zip,windows,amd64,.exe)