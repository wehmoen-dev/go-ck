package types

type Client interface {
	GetRecipe(recipeId string) (*Recipe, error)
	GetComments(recipeId string) ([]*Comment, error)
	GetImages(recipeId string) ([]*RecipeImage, error)
	GetUser(userId string, includeProfile bool) (*User, error)
	SearchRecipes(query string, limit int) ([]*Recipe, error)
}
