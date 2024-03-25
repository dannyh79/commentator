package main

import (
	repo "github.com/dannyh79/commentator/internal/repo"
	routes "github.com/dannyh79/commentator/internal/rest/v1"
	shows "github.com/dannyh79/commentator/internal/shows"
	"github.com/gin-gonic/gin"
)

// ENHANCEMENT: Handle os.signal to gracefully shut down app
func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	// ENHANCEMENT: Use repo that persists data
	r := repo.NewInMemoryCommentRepo()
	u := shows.NewShowComments(r)

	routes.AddRoutes(engine, u)
	engine.Run()
}
