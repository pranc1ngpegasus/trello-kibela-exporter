package usecase

import (
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"

	"github.com/adlio/trello"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ ImportTrello = (*importTrello)(nil)

type (
	ImportTrello interface {
		Do(input ImportTrelloInput) (*ImportTrelloOutput, error)
	}

	importTrello struct {
		config       configuration.Config
		logger       *logrus.Logger
		trelloClient *trello.Client
	}
)

func NewImportTrello(
	config configuration.Config,
	logger *logrus.Logger,
) ImportTrello {
	return &importTrello{
		config: config,
		logger: logger,
		trelloClient: newTrelloClient(
			config,
			logger,
		),
	}
}

func newTrelloClient(
	config configuration.Config,
	logger *logrus.Logger,
) *trello.Client {
	trelloClient := trello.NewClient(
		config.Trello.APIKey,
		config.Trello.Token,
	)
	trelloClient.Logger = logger

	return trelloClient
}

type (
	ImportTrelloInput struct {
		BoardID string
	}

	ImportTrelloOutput struct {
		Board *trello.Board
	}
)

var (
	ErrBoardNotFound error = errors.New("board not found")
	ErrListsNotFound error = errors.New("lists not found")
)

func (u *importTrello) Do(input ImportTrelloInput) (*ImportTrelloOutput, error) {
	board, err := u.trelloClient.GetBoard(input.BoardID, trello.Defaults())
	if err != nil {
		return nil, err
	}
	if board == nil {
		return nil, ErrBoardNotFound
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, err
	}
	if len(lists) == 0 {
		return nil, ErrListsNotFound
	}

	board.Lists = lists

	for _, list := range lists {
		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			return nil, err
		}
		if len(cards) == 0 {
			continue
		}

		list.Cards = cards
	}

	return &ImportTrelloOutput{
		Board: board,
	}, nil
}
