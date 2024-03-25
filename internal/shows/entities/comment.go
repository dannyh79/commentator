package entities

type Comment struct {
	Id      int
	UserId  int
	ShowId  int
	Comment string
	Upvote  int
}

func NewComment(d struct {
	Id      int
	UserId  int
	ShowId  int
	Comment string
	Upvote  int
}) *Comment {
	return &Comment{
		Id:      d.Id,
		UserId:  d.UserId,
		ShowId:  d.ShowId,
		Comment: d.Comment,
		Upvote:  0,
	}
}
