package main

import (
	"fmt"

	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
)

func init() {
	configuration.Load()
}

func main() {
	cmd := initialize()
	kibela, err := cmd.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(kibela.NoteID)
}
