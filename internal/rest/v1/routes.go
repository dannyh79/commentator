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

type FailureOutput struct {
	Error string `json:"error"`
}

func AddRoutes(r *gin.Engine, u *shows.ShowComments) {
	v1 := r.Group("/v1")
	v1.POST("/shows/:id/comments", createCommentHandler(u))
}

func createCommentHandler(u *shows.ShowComments) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload PostShowsCommentsInput
		c.ShouldBindJSON(&payload)
		showId, _ := strconv.Atoi(c.Param("id"))
		params := toCreateCommentParams(showId, &payload)
		comment, err := u.CreateComment(params)
		if err != nil {
			c.JSON(http.StatusBadRequest, FailureOutput{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, toPostShowsCommentsOutput(comment))
	}
}

func toCreateCommentParams(showId int, p *PostShowsCommentsInput) *shows.CreateCommentInput {
	return &shows.CreateCommentInput{
		UserId:  p.UserId,
		ShowId:  showId,
		Comment: p.Comment,
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
