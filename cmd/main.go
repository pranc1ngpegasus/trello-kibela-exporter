package main

import (
	"fmt"

	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/handler"
)

func init() {
	configuration.Load()
}

func main() {
	cmd := initialize()
	kibela, err := cmd.Do(
		handler.TrelloToKibelaInput{
			BoardID: "xxx",
			Folder:  "xxx",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(kibela.NoteID)
}
