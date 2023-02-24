#!/bin/bash


run: main.go logger.go
	go run $^

gen_denpendence:
	go mod tidy

install-dependence:
	go mod download

import-dependence:
	go mod vendor
