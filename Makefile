.PHONY: build

build:
	go build -ldflags "-X github.com/yoyrandao/osctx/cmd.Version=1.0.0 -X github.com/yoyrandao/osctx/cmd.Commit=af71b56" -o dist/osctx main.go

install: build
	sudo cp dist/osctx /usr/local/bin/