.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot ~/Bot-project/main.go
run: build
	./.bin/bot