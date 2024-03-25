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

var ErrorEmptyComment = errors.New("empty comment found")

var ErrorInValidComment = errors.New("invalid word found in comment")

func (u *ShowComments) CreateComment(i *CreateCommentInput) (*CreateCommentOutput, error) {
	// ENHANCEMENT: validate showId by checking if such show exists. Returns error if not.

	c := newComment(i)

	// ENHANCEMENT: run validations parallelly if they are costy
	if isEmptyComment(c.Comment) {
		return nil, ErrorEmptyComment
	}
	if isInvalidComment(c.Comment) {
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

func isEmptyComment(c string) bool {
	if len(strings.TrimSpace(c)) == 0 {
		return true
	}
	return false
}

// Arbitrary list of strings that are deemed invalid
var InvalidWords = []string{
	"NSFW",
}

func isInvalidComment(c string) bool {
	for _, word := range InvalidWords {
		if strings.Contains(c, word) {
			return true
		}
	}
	return false
}
