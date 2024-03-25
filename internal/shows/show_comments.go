package shows

type ShowComments struct{}

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

func NewShowComments() *ShowComments {
	return &ShowComments{}
}
