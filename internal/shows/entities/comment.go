package entities

// ENHANCEMENT: set default field value
// ENHANCEMENT: leverage UUID to prevent IDOR (insecure direct object references)
type Comment struct {
	Id      int
	UserId  int // foreign key
	ShowId  int // foreign key
	Comment string
	Upvote  int
}
