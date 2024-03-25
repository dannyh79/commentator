package repo

import "github.com/dannyh79/commentator/internal/shows/entities"

type Repository[T any] interface {
	// ENHANCEMENT: use (data, err) signature support error handling on consumer side
	Save(*T) T
}

// ENHANCEMENT: set default field value
// ENHANCEMENT: leverage UUID to prevent IDOR (insecure direct object references)
type CommentSchema struct {
	Id      int
	UserId  int
	ShowId  int
	Comment string
	Upvote  int
}

type InMemoryCommmentRepo struct {
	data map[int]CommentSchema
}

func (r *InMemoryCommmentRepo) Save(c *entities.Comment) entities.Comment {
	r.data[c.Id] = *toCommentSchema(c)
	row := r.data[c.Id]
	return *toComment(&row)
}

func NewInMemoryCommentRepo() *InMemoryCommmentRepo {
	return &InMemoryCommmentRepo{
		data: map[int]CommentSchema{},
	}
}

func toComment(c *CommentSchema) *entities.Comment {
	return &entities.Comment{
		Id:      c.Id,
		UserId:  c.UserId,
		ShowId:  c.Id,
		Comment: c.Comment,
		Upvote:  c.Upvote,
	}
}

func toCommentSchema(c *entities.Comment) *CommentSchema {
	return &CommentSchema{
		Id:      c.Id,
		UserId:  c.UserId,
		ShowId:  c.Id,
		Comment: c.Comment,
		Upvote:  c.Upvote,
	}
}
