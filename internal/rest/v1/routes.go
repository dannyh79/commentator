package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatedComment struct {
	UserId  int    `json:"user_id"`
	Comment string `json:"comment"`
	Upvote  int    `json:"upvote"`
}
type PostShowsCommentsOutput struct {
	Result CreatedComment `json:"result"`
}

func AddRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.POST("/shows/:id/comments", func(c *gin.Context) {
		c.JSON(http.StatusCreated, toPostShowsCommentsOutput())
	})
}

func toPostShowsCommentsOutput() *PostShowsCommentsOutput {
	return &PostShowsCommentsOutput{
		Result: CreatedComment{
			UserId:  2,
			Comment: "some comments",
			Upvote:  0,
		},
	}
}
