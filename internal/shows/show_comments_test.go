package shows_test

import (
	"testing"

	"github.com/dannyh79/commentator/internal/repo"
	shows "github.com/dannyh79/commentator/internal/shows"
	utils "github.com/dannyh79/commentator/internal/testutils"
)

func Test_ShowComments_CreateComment(t *testing.T) {
	tests := []struct {
		name          string
		params        *shows.CreateCommentInput
		expected      *shows.CreateCommentOutput
		expectError   bool
		expectedError error
	}{
		{
			name:        "creates a comment",
			params:      &shows.CreateCommentInput{UserId: 2, ShowId: 1, Comment: "some comments"},
			expected:    &shows.CreateCommentOutput{UserId: 2, Comment: "some comments", Upvote: 0},
			expectError: false,
		},
		{
			name:          "returns error ErrorEmptyComment",
			params:        &shows.CreateCommentInput{UserId: 2, ShowId: 1, Comment: "  "},
			expected:      nil,
			expectError:   true,
			expectedError: shows.ErrorEmptyComment,
		},
		{
			name:          "returns error ErrorInValidComment",
			params:        &shows.CreateCommentInput{UserId: 2, ShowId: 1, Comment: "some NSFW comments"},
			expected:      nil,
			expectError:   true,
			expectedError: shows.ErrorInValidComment,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := repo.NewInMemoryCommentRepo()
			usecase := shows.NewShowComments(r)

			got, err := usecase.CreateComment(tc.params)

			if tc.expectError {
				utils.AssertErrorEqual(t)(err, tc.expectedError)
			}
			utils.AssertEqual(t)(got, tc.expected)
		})
	}
}
