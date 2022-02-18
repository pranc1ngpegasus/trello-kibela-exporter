package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/adlio/trello"
	"github.com/sirupsen/logrus"
)

var _ ConstructMarkdown = (*constructMarkdown)(nil)

type (
	ConstructMarkdown interface {
		Do(input ConstructMarkdownInput) (*ConstructMarkdownOutput, error)
	}

	constructMarkdown struct {
		logger *logrus.Logger
	}
)

func NewConstructMarkdown(
	logger *logrus.Logger,
) ConstructMarkdown {
	return &constructMarkdown{
		logger: logger,
	}
}

type (
	ConstructMarkdownInput struct {
		TrelloBoard  *trello.Board
		BoardMembers map[string]BoardMember
	}

	ConstructMarkdownOutput struct {
		Title   string
		Content string
	}
)

func (u *constructMarkdown) Do(input ConstructMarkdownInput) (*ConstructMarkdownOutput, error) {
	now := time.Now()
	nowDay := now.Format("20060102")

	// Describe MD title
	title := fmt.Sprintf("%s_%s", nowDay, input.TrelloBoard.Name)

	var contentRows []string
	for _, list := range input.TrelloBoard.Lists {
		if list == nil {
			continue
		}

		listRow, err := u.constructList(list, input.BoardMembers)
		if err != nil {
			u.logger.Error(err)
			return nil, err
		}

		contentRows = append(contentRows, listRow)
	}

	return &ConstructMarkdownOutput{
		Title:   title,
		Content: strings.Join(contentRows, "\n"),
	}, nil
}

func (u *constructMarkdown) constructList(list *trello.List, members map[string]BoardMember) (string, error) {
	var sectionRows []string

	// section title
	sectionRows = append(sectionRows, fmt.Sprintf("## %s", list.Name))

	for _, card := range list.Cards {
		if card == nil {
			continue
		}

		cardRow, err := u.constructCard(card, members)
		if err != nil {
			u.logger.Error(err)
			return "", err
		}

		sectionRows = append(sectionRows, cardRow)
	}

	return strings.Join(sectionRows, "\n"), nil
}

func (u *constructMarkdown) constructCard(card *trello.Card, members map[string]BoardMember) (string, error) {
	var sectionRows []string
	var sectionRow string

	sectionRow = sectionRow + fmt.Sprintf("### %s", card.Name)
	if len(card.IDMembers) > 0 {
		sectionRow = sectionRow + fmt.Sprintf(" by %s", u.memberNameByID(
			card.IDMembers[0],
			members,
		))
	}
	// section title
	sectionRows = append(sectionRows, sectionRow)

	// section content
	sectionRows = append(sectionRows, fmt.Sprintf("%s", card.Desc))

	return strings.Join(sectionRows, "\n"), nil
}

func (u *constructMarkdown) memberNameByID(id string, members map[string]BoardMember) string {
	member, exist := members[id]
	if !exist {
		return ""
	}

	return member.FullName
}
