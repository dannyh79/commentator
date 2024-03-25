package entities

type Comment struct {
	Id      int
	UserId  int
	ShowId  int
	Comment string
	Upvote  int
}
