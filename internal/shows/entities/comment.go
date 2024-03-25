package entities

// ENHANCEMENT: set default field value
// ENHANCEMENT: leverage UUID to prevent IDOR (insecure direct object references)
type Comment struct {
	Id      int
	UserId  int
	ShowId  int
	Comment string
	Upvote  int
}
