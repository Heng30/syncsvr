#!/bin/bash

APP_NAME=syncsvr
DATE=`date "+%Y_%m_%d"`
VERSION=`git tag | tail -n 1`

run: main.go version.go
	go run $^

build:
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	go build -o ${APP_NAME}_${VERSION}_${DATE}

build-realse:
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	go build -ldflags "-s -w" -o ${APP_NAME}_${VERSION}_${DATE}

build-arm:
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ${APP_NAME}_arm_${VERSION}_${DATE}

gen_denpendence:
	go mod tidy

download-dependence:
	go mod download

import-dependence:
	go mod vendor



ping:
	curl "http://localhost:8000/ping"
