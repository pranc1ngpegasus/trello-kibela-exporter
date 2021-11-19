package usecase

import (
	"strings"

	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"

	"github.com/adlio/trello"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ ArchiveTrello = (*archiveTrello)(nil)

type (
	ArchiveTrello interface {
		Do(input ArchiveTrelloInput) error
	}

	archiveTrello struct {
		config       configuration.Config
		logger       *logrus.Logger
		trelloClient *trello.Client
	}
)

func NewArchiveTrello(
	config configuration.Config,
	logger *logrus.Logger,
) ArchiveTrello {
	return &archiveTrello{
		config: config,
		logger: logger,
		trelloClient: newTrelloClient(
			config,
			logger,
		),
	}
}

type (
	ArchiveTrelloInput struct {
		BoardID     string
		IgnoreLists []string
	}
)

func (u *archiveTrello) Do(input ArchiveTrelloInput) error {
	board, err := u.trelloClient.GetBoard(input.BoardID, trello.Defaults())
	if err != nil {
		return err
	}
	if board == nil {
		return errors.New("")
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return err
	}
	if len(lists) == 0 {
		return ErrListsNotFound
	}

	board.Lists = lists

	for _, list := range lists {
		if u.isIgnored(list, input.IgnoreLists) {
			continue
		}

		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			return err
		}
		if len(cards) == 0 {
			continue
		}

		for _, card := range cards {
			if err := card.Archive(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *archiveTrello) isIgnored(list *trello.List, ignoreList []string) bool {
	if list == nil {
		return false
	}

	for _, ignoreText := range ignoreList {
		if strings.Contains(list.Name, ignoreText) {
			return true
		}
	}

	return false
}
