package main

import (
	"encoding/json"
	"fmt"

	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/usecase"
)

const (
	testContent = "# これはなに\nほげほげ"
)

func init() {
	configuration.Load()
}

func main() {
	cmd := initialize()
	board, err := cmd.Do(usecase.ExportKibelaInput{
		Title:   "trello-kibela-exporterのテスト",
		Content: testContent,
		CoEdit:  true,
		Folder:  "xxx",
		Groups:  []string{"xxx"},
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
