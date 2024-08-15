package test

import "testing"

func TestGetImages(t *testing.T) {
	images, err := c.GetImages(TestRecipeId)
	if err != nil {
		t.Fatalf("Failed to get images: %v", err)
	}

	if len(images) == 0 {
		t.Fatalf("No images found. There should be images for the test recipe.")
	}
}
