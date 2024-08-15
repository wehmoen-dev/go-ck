package types

import "strings"

type RecipeImage struct {
	Id          string      `json:"id"`
	Rating      Rating      `json:"rating"`
	Owner       User        `json:"owner"`
	Credits     interface{} `json:"credits"`
	UrlTemplate string      `json:"urlTemplate,omitempty"`
	Url         string      `json:"url,omitempty"`
}

func (r *RecipeImage) Parse() {
	r.Url = strings.ReplaceAll(r.UrlTemplate, "<format>", CdnImageFormat)
	r.UrlTemplate = ""
}
