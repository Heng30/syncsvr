#!/bin/bash

run: main.go
	go run $^

build:
	go build

gen_denpendence:
	go mod tidy

download-dependence:
	go mod download

import-dependence:
	go mod vendor



ping:
	curl "http://localhost:8080/ping"

markCoins:
	curl -v -H "Origin: heng30.com" "http://localhost:8080/1234/markCoins"
