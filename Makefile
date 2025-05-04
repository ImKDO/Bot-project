.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot ~/Bot-project/cmd/main.go
run: build
	./.bin/bot