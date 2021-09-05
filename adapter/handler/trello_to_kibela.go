package handler

import (
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/usecase"
	"github.com/sirupsen/logrus"
)

var _ TrelloToKibela = (*trelloToKibela)(nil)

type (
	TrelloToKibela interface {
		Do(input TrelloToKibelaInput) (*TrelloToKibelaOutput, error)
	}

	trelloToKibela struct {
		config                   configuration.Config
		logger                   *logrus.Logger
		importTrelloUsecase      usecase.ImportTrello
		getBoardMembersUsecase   usecase.GetBoardMembers
		constructMarkdownUsecase usecase.ConstructMarkdown
		exportKibela             usecase.ExportKibela
	}
)

func NewTrelloToKibela(
	config configuration.Config,
	logger *logrus.Logger,
	importTrelloUsecase usecase.ImportTrello,
	getBoardMembersUsecase usecase.GetBoardMembers,
	constructMarkdownUsecase usecase.ConstructMarkdown,
	exportKibela usecase.ExportKibela,
) TrelloToKibela {
	return &trelloToKibela{
		config:                   config,
		logger:                   logger,
		importTrelloUsecase:      importTrelloUsecase,
		getBoardMembersUsecase:   getBoardMembersUsecase,
		constructMarkdownUsecase: constructMarkdownUsecase,
		exportKibela:             exportKibela,
	}
}

type (
	TrelloToKibelaInput struct {
		BoardID string
		Folder  string
	}

	TrelloToKibelaOutput struct {
		NoteID string
	}
)

func (h *trelloToKibela) Do(input TrelloToKibelaInput) (*TrelloToKibelaOutput, error) {
	trelloBoard, err := h.importTrelloUsecase.Do(
		usecase.ImportTrelloInput{
			BoardID: input.BoardID,
		},
	)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	h.logger.Infof("%+v", *trelloBoard)

	boardMembers, err := h.getBoardMembersUsecase.Do(
		usecase.GetBoardMembersInput{
			BoardID: input.BoardID,
		},
	)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	markdown, err := h.constructMarkdownUsecase.Do(
		usecase.ConstructMarkdownInput{
			TrelloBoard:  trelloBoard.Board,
			BoardMembers: boardMembers.Members,
		},
	)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	h.logger.Infof("%+v", markdown)

	kibela, err := h.exportKibela.Do(
		usecase.ExportKibelaInput{
			Title:   markdown.Title,
			Content: markdown.Content,
			CoEdit:  h.config.Kibela.CoEdit,
			Folder:  input.Folder,
			Groups: []string{
				h.config.Kibela.Group,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	h.logger.Infof("%+v", kibela)

	return &TrelloToKibelaOutput{
		NoteID: kibela.NoteID,
	}, nil
}
