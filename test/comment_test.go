package test

import "testing"

func TestGetComments(t *testing.T) {
	comments, err := c.GetComments(TestRecipeId)
	if err != nil {
		t.Fatalf("Failed to get comments: %v", err)
	}

	if len(comments) == 0 {
		t.Fatalf("No comments found. There should be comments for the test recipe.")
	}
}
