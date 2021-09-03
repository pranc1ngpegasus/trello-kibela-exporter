package main

import (
	"encoding/json"
	"fmt"

	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/usecase"
)

func init() {
	configuration.Load()
}

func main() {
	cmd := initialize()
	board, err := cmd.Do(usecase.ImportTrelloInput{
		BoardID: "xxxxx",
	})
	if err != nil {
		panic(err)
	}

	e, err := json.Marshal(board)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(e))
}
