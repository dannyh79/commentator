package shows_test

import (
	"testing"

	shows "github.com/dannyh79/commentator/internal/shows"
	utils "github.com/dannyh79/commentator/internal/testutils"
)

func Test_ShowComments_CreateComment(t *testing.T) {
	tests := []struct {
		name     string
		params   *shows.CreateCommentInput
		expected *shows.CreateCommentOutput
	}{
		{
			name:     "creates a comment",
			params:   &shows.CreateCommentInput{UserId: 2, ShowId: 1, Comment: "some comments"},
			expected: &shows.CreateCommentOutput{UserId: 2, Comment: "some comments", Upvote: 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			usecase := shows.NewShowComments()

			got := usecase.CreateComment(tc.params)

			utils.AssertEqual(t)(got, tc.expected)
		})
	}
}
