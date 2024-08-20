package client

import (
	"fmt"
	"github.com/wehmoen-dev/go-ck/pkg/types"
)

func (c *c) GetRecipe(recipeId string) (*types.Recipe, error) {
	var result *types.Recipe

	response, err := c.restClient.R().SetResult(&result).Get(fmt.Sprintf("/recipes/%s", recipeId))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("received non 2xx response: %d - %v", response.StatusCode(), response)
	}

	result.Parse()
	result.SetRestClient(c.restClient)

	return result, nil
}
