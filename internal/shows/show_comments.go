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

func (u *ShowComments) CreateComment(i *CreateCommentInput) *CreateCommentOutput {
	c := newComment(i)
	r := u.repo.Save(c)
	return &CreateCommentOutput{
		UserId:  r.UserId,
		Comment: r.Comment,
		Upvote:  r.Upvote,
	}
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
