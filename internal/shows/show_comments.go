package shows

import (
	"errors"
	"strings"

	"github.com/dannyh79/commentator/internal/repo"
	"github.com/dannyh79/commentator/internal/shows/entities"
)

type ShowComments struct {
	repo repo.Repository[entities.Comment]
}

type CreateCommentInput struct {
	UserId  int
	ShowId  int
	Comment string
}
type CreateCommentOutput struct {
	UserId  int
	Comment string
	Upvote  int
}

var ErrorInValidComment = errors.New("invalid word found in comment")

func (u *ShowComments) CreateComment(i *CreateCommentInput) (*CreateCommentOutput, error) {
	c := newComment(i)
	if !isValidComment(c.Comment) {
		return nil, ErrorInValidComment
	}

	r := u.repo.Save(c)
	return &CreateCommentOutput{
		UserId:  r.UserId,
		Comment: r.Comment,
		Upvote:  r.Upvote,
	}, nil
}

func NewShowComments(r repo.Repository[entities.Comment]) *ShowComments {
	return &ShowComments{r}
}

func newComment(i *CreateCommentInput) *entities.Comment {
	return &entities.Comment{
		UserId:  i.UserId,
		ShowId:  i.ShowId,
		Comment: i.Comment,
		Upvote:  0,
	}
}

// Arbitrary list of strings that are deemed invalid
var InValidWords = []string{
	"NSFW",
}

func isValidComment(c string) bool {
	for _, word := range InValidWords {
		if strings.Contains(c, word) {
			return false
		}
	}
	return true
}
