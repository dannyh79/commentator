package shows

import "github.com/dannyh79/commentator/internal/repo"

type ShowComments struct {
	repo repo.Repository[repo.Comment]
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

func NewShowComments(r repo.Repository[repo.Comment]) *ShowComments {
	return &ShowComments{r}
}
