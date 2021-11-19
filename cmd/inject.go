//+build wireinject

package main

import (
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/handler"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/logger"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/usecase"

	"github.com/google/wire"
)

func initialize() handler.TrelloToKibela {
	wire.Build(
		configuration.Get,
		logger.Logger,
		handler.NewTrelloToKibela,
		usecase.NewImportTrello,
		usecase.NewGetBoardMembers,
		usecase.NewConstructMarkdown,
		usecase.NewExportKibela,
		usecase.NewArchiveTrello,
	)

	return nil
}
