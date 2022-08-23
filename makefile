SHELL := /bin/bash

# go install github.com/divan/expvarmon@latest

run:
	go run app/ingress-check/main.go

tidy:
	go mod tidy
	go mod vendor

docker:
	docker buildx build --platform=linux/amd64,linux/arm64 -t wonko/ingress-check:0.0.1 -t wonko/ingress-check:latest --push .
