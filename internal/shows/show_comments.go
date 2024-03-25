package shows

import (
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

func (u *ShowComments) CreateComment(*CreateCommentInput) *CreateCommentOutput {
	return &CreateCommentOutput{
		UserId:  2,
		Comment: "some comments",
		Upvote:  0,
	}
}

func NewShowComments(r repo.Repository[entities.Comment]) *ShowComments {
	return &ShowComments{r}
}
