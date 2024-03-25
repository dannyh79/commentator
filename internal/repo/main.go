package repo

type Repository[T any] interface {
	Save(*T) T
}

type Comment struct {
	Id      int
	UserId  int
	ShowId  int
	Comment string
	Upvote  int
}

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

func (r *InMemoryCommmentRepo) Save(c *Comment) Comment {
	r.data[c.Id] = *toCommentSchema(c)
	row := r.data[c.Id]
	return *toComment(&row)
}

func NewInMemoryCommentRepo() *InMemoryCommmentRepo {
	return &InMemoryCommmentRepo{
		data: map[int]CommentSchema{},
	}
}

func toComment(c *CommentSchema) *Comment {
	return &Comment{
		Id:      c.Id,
		UserId:  c.UserId,
		ShowId:  c.Id,
		Comment: c.Comment,
		Upvote:  c.Upvote,
	}
}

func toCommentSchema(c *Comment) *CommentSchema {
	return &CommentSchema{
		Id:      c.Id,
		UserId:  c.UserId,
		ShowId:  c.Id,
		Comment: c.Comment,
		Upvote:  c.Upvote,
	}
}
