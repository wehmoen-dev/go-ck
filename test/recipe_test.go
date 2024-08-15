package test

import (
	"testing"
)

func TestGetRecipe(t *testing.T) {
	recipe, err := c.GetRecipe(TestRecipeId)
	if err != nil {
		t.Fatalf("Failed to get recipe: %v", err)
	}

	if recipe == nil {
		t.Fatalf("Recipe is nil")
	}
}

func TestGetRecipeImages(t *testing.T) {
	recipe, err := c.GetRecipe(TestRecipeId)
	if err != nil {
		t.Fatalf("Failed to get recipe: %v", err)
	}

	if recipe == nil {
		t.Fatalf("Recipe is nil")
	}

	images, err := recipe.GetImages()

	if err != nil {
		t.Fatalf("Failed to get images: %v", err)
	}

	if len(images) == 0 {
		t.Fatalf("No images found. There should be images for the test recipe.")
	}

}
