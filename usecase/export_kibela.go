package usecase

import (
	"context"
	"strings"

	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"

	"github.com/Songmu/kibelasync/kibela"
	"github.com/sirupsen/logrus"
)

var _ ExportKibela = (*exportKibela)(nil)

type (
	ExportKibela interface {
		Do(input ExportKibelaInput) error
	}

	exportKibela struct {
		config       configuration.Config
		logger       *logrus.Logger
		kibelaClient *kibela.Kibela
	}
)

func NewExportKibela(
	config configuration.Config,
	logger *logrus.Logger,
) ExportKibela {
	return &exportKibela{
		config:       config,
		logger:       logger,
		kibelaClient: mustNewKibelaClient(config),
	}
}

func newKibelaClient(
	config configuration.Config,
) (*kibela.Kibela, error) {
	return kibela.New("trello-kibela-exporter")
}

func mustNewKibelaClient(
	config configuration.Config,
) *kibela.Kibela {
	client, err := newKibelaClient(config)
	if err != nil {
		panic(err)
	}

	return client
}

type (
	ExportKibelaInput struct {
		Title   string
		Content string
		CoEdit  bool
		Folder  string
		Groups  []string
	}
)

func (u *exportKibela) Do(input ExportKibelaInput) error {
	md, err := kibela.NewMD(
		"",
		strings.NewReader(input.Content),
		input.Title,
		input.CoEdit,
		"",
	)
	if err != nil {
		u.logger.Error(err)
		return err
	}

	if input.Content != "" {
		md.Content = input.Content
	}

	if input.Title != "" {
		md.FrontMatter.Title = input.Title
	}

	if input.Folder != "" {
		md.FrontMatter.Folder = input.Folder
	}

	if len(input.Groups) > 0 {
		md.FrontMatter.Groups = input.Groups
	}

	if err := u.kibelaClient.PublishMD(
		context.Background(),
		md,
		false,
	); err != nil {
		u.logger.Error(err)
		return err
	}

	return nil
}
