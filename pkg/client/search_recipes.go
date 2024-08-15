package client

import (
	"github.com/wehmoen/go-ck/pkg/types"
	"strconv"
)

func (c *c) SearchRecipes(query string) ([]*types.Recipe, error) {
	return c.searchRecipes(query, 0)
}

func (c *c) searchRecipes(query string, offset int) ([]*types.Recipe, error) {
	var recipes []*types.Recipe

	for {
		var result struct {
			Count   int `json:"count"`
			Results []struct {
				Recipe *types.Recipe `json:"recipe"`
				Score  int           `json:"score"`
			} `json:"results"`
		}

		_, err := c.restClient.R().
			SetResult(&result).
			SetQueryParamsFromValues(map[string][]string{
				"query":             {query},
				"limit":             {strconv.Itoa(types.PaginationDefaultLimit)},
				"offset":            {strconv.Itoa(offset)},
				"orderBy":           {"2"},
				"descendCategories": {"1"},
				"order":             {"0"},
			}).
			Get("/recipes")

		if err != nil {
			return nil, err
		}

		for _, recipe := range result.Results {
			recipes = append(recipes, recipe.Recipe)
		}

		if len(result.Results) < types.PaginationDefaultLimit {
			break
		}
		offset += len(result.Results)
	}

	return recipes, nil
}
