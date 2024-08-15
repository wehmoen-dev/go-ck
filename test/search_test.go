package test

import "testing"

func TestSearchRecipes(t *testing.T) {
	recipes, err := c.SearchRecipes(TestSearchQuery)
	if err != nil {
		t.Fatalf("Failed to search recipes: %v", err)
	}

	if len(recipes) == 0 {
		t.Fatalf("No recipes found. There should be recipes for the test search query.")
	}
}
