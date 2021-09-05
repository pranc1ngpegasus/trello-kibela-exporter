//+build wireinject

package main

import (
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/logger"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/usecase"

	"github.com/google/wire"
)

func initialize() usecase.ExportKibela {
	wire.Build(
		configuration.Get,
		logger.Logger,
		usecase.NewExportKibela,
	)

	return nil
}
