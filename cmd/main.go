package main

import (
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
)

func init() {
	configuration.Load()
}

func main() {
	cmd := initialize()
	if err := cmd.Do(); err != nil {
		panic(err)
	}
}
