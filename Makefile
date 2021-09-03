.PHONY: prepare-tools wiregen

prepare-tools:
	go get github.com/google/wire/cmd/wire

wiregen: prepare-tools
	go run github.com/google/wire/cmd/wire gen ./...
