.PHONY: manifest
manifest:
	kustomize build .

.PHONY: build
build:
	go build -o ./lbtest .

.PHONY: run
run: build
	INTERVAL=1s TIMEOUT=1s TARGET=http://google.com ./lbtest
