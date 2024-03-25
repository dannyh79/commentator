package routes

import (
	"net/http"
	"strconv"

	"github.com/dannyh79/commentator/internal/shows"
	"github.com/gin-gonic/gin"
)

type PostShowsCommentsInput struct {
	UserId  int    `json:"user_id"`
	Comment string `json:"comment"`
}
type CreatedComment struct {
	UserId  int    `json:"user_id"`
	Comment string `json:"comment"`
	Upvote  int    `json:"upvote"`
}
type PostShowsCommentsOutput struct {
	Result CreatedComment `json:"result"`
}

func AddRoutes(r *gin.Engine, u *shows.ShowComments) {
	v1 := r.Group("/v1")
	v1.POST("/shows/:id/comments", createCommentHandler(u))
}

func createCommentHandler(u *shows.ShowComments) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload shows.CreateCommentInput
		c.ShouldBindJSON(&payload)
		showId, _ := strconv.Atoi(c.Param("id"))
		payload.ShowId = showId
		comment := u.CreateComment(&payload)
		c.JSON(http.StatusCreated, toPostShowsCommentsOutput(comment))
	}
}

func toPostShowsCommentsOutput(c *shows.CreateCommentOutput) *PostShowsCommentsOutput {
	return &PostShowsCommentsOutput{
		Result: CreatedComment{
			UserId:  c.UserId,
			Comment: c.Comment,
			Upvote:  c.Upvote,
		},
	}
}
