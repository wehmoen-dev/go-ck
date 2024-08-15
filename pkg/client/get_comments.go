package client

import (
	"github.com/wehmoen/go-ck/pkg/types"
)

func (c *c) GetComments(recipeId string) ([]*types.Comment, error) {
	recipe, err := c.GetRecipe(recipeId)
	if err != nil {
		return nil, err
	}
	return recipe.GetComments()
}
