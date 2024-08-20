package client

import (
	"github.com/wehmoen-dev/go-ck/pkg/types"
)

func (c *c) GetImages(recipeId string) ([]*types.RecipeImage, error) {
	recipe, err := c.GetRecipe(recipeId)
	if err != nil {
		return nil, err
	}
	return recipe.GetImages()
}
