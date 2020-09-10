cmd = ./cmd/go-hh-parser/
output = ./bin/
log = ./log

all = build run
.PHONY: all

build:
	mkdir -p $(output)
	go build -v -o $(output) $(cmd)

run:
	mkdir -p $(log)
	go run -v $(cmd) &> ${log}/log.log

.DEFAULT_GOAL := build
