package types

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

type Rating struct {
	Rating          float64 `json:"rating"`
	NumberOfRatings int     `json:"numVotes"`
}

type IngredientGroup struct {
	Header      string       `json:"header"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Id                      string  `json:"id"`
	Name                    string  `json:"name"`
	Unit                    string  `json:"unit"`
	Amount                  float64 `json:"amount"`
	UsageInfo               string  `json:"usageInfo"`
	ProductGroup            string  `json:"productGroup"`
	PreviewImageUrlTemplate string  `json:"previewImageUrlTemplate,omitempty"`
	PreviewImageUrl         string  `json:"previewImageUrl,omitempty"`
}

type Recipe struct {
	restClient *resty.Client

	Id                      string            `json:"id"`
	Title                   string            `json:"title"`
	Subtitle                string            `json:"subtitle"`
	Rating                  Rating            `json:"rating"`
	Difficulty              int               `json:"difficulty"`
	HasImage                bool              `json:"hasImage"`
	HasVideo                bool              `json:"hasVideo"`
	PreviewImageUrlTemplate string            `json:"previewImageUrlTemplate,omitempty"`
	PreviewImageUrl         string            `json:"previewImageUrl,omitempty"`
	ImageCount              int               `json:"imageCount"`
	IsPlus                  bool              `json:"isPlus"`
	ViewCount               int               `json:"viewCount"`
	Servings                int               `json:"servings"`
	CookingTime             int               `json:"cookingTime"`
	RestingTime             int               `json:"restingTime"`
	TotalTime               int               `json:"totalTime"`
	SiteUrl                 string            `json:"siteUrl"`
	Instructions            string            `json:"instructions"`
	IngredientGroups        []IngredientGroup `json:"ingredientGroups"`
	APIUrl                  string            `json:"apiUrl,omitempty"`
}

func (r *Recipe) SetRestClient(client *resty.Client) {
	r.restClient = client
}

func (r *Recipe) Parse() {
	var groups []IngredientGroup

	for _, group := range r.IngredientGroups {

		var newGroup IngredientGroup
		newGroup.Header = group.Header
		newGroup.Ingredients = []Ingredient{}

		for _, item := range group.Ingredients {
			item.PreviewImageUrl = strings.ReplaceAll(item.PreviewImageUrlTemplate, "<format>", CdnImageFormat)
			item.PreviewImageUrlTemplate = ""
			newGroup.Ingredients = append(newGroup.Ingredients, item)
		}
		groups = append(groups, newGroup)
	}

	r.IngredientGroups = groups
}

func (r *Recipe) GetImages() ([]*RecipeImage, error) {
	if r.restClient == nil {
		return nil, fmt.Errorf("rest client not set")
	}
	return r.getImages(0)
}

func (r *Recipe) GetComments() ([]*Comment, error) {
	if r.restClient == nil {
		return nil, fmt.Errorf("rest client not set")
	}
	return r.getComments(0)
}

func (r *Recipe) getComments(offset int) ([]*Comment, error) {
	var comments []*Comment

	for {

		var result struct {
			Count   int `json:"count"`
			Results []*Comment
		}

		_, err := r.restClient.R().SetResult(&result).Get(fmt.Sprintf("/recipes/%s/comments?offset=%d&limit=%d", r.Id, offset, PaginationDefaultLimit))

		if err != nil {
			return nil, err
		}

		comments = append(comments, result.Results...)

		if len(result.Results) < PaginationDefaultLimit {
			break
		}
		offset += len(result.Results)
	}

	return comments, nil
}

func (r *Recipe) getImages(offset int) ([]*RecipeImage, error) {
	var images []*RecipeImage

	for {

		var result struct {
			Count   int `json:"count"`
			Results []*RecipeImage
		}

		_, err := r.restClient.R().SetResult(&result).Get(fmt.Sprintf("/recipes/%s/images?offset=%d&limit=%d", r.Id, offset, PaginationDefaultLimit))

		if err != nil {
			return nil, err
		}

		for _, image := range result.Results {
			image.Parse()
			images = append(images, image)
		}

		if len(result.Results) < PaginationDefaultLimit {
			break
		}
		offset += len(result.Results)
	}

	return images, nil
}
