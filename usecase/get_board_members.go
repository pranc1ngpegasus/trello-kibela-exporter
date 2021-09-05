package usecase

import (
	"github.com/Pranc1ngPegasus/trello-kibela-exporter/adapter/configuration"

	"github.com/adlio/trello"
	"github.com/sirupsen/logrus"
)

type (
	GetBoardMembers interface {
		Do(input GetBoardMembersInput) (*GetBoardMembersOutput, error)
	}

	getBoardMembers struct {
		config       configuration.Config
		logger       *logrus.Logger
		trelloClient *trello.Client
	}
)

func NewGetBoardMembers(
	config configuration.Config,
	logger *logrus.Logger,
) GetBoardMembers {
	return &getBoardMembers{
		config: config,
		logger: logger,
		trelloClient: newTrelloClient(
			config,
			logger,
		),
	}
}

type (
	GetBoardMembersInput struct {
		BoardID string
	}

	GetBoardMembersOutput struct {
		Members map[string]BoardMember
	}

	BoardMember struct {
		ID       string
		FullName string
	}
)

func (u *getBoardMembers) Do(input GetBoardMembersInput) (*GetBoardMembersOutput, error) {
	board, err := u.trelloClient.GetBoard(input.BoardID, trello.Defaults())
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}

	members, err := board.GetMembers(trello.Defaults())
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}

	boardMembers := make(map[string]BoardMember)
	for _, member := range members {
		boardMembers[member.ID] = BoardMember{
			ID:       member.ID,
			FullName: member.FullName,
		}
	}

	return &GetBoardMembersOutput{
		Members: boardMembers,
	}, nil
}
