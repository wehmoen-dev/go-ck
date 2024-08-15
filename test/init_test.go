package test

import (
	"github.com/wehmoen/go-ck/pkg/client"
	"github.com/wehmoen/go-ck/pkg/types"
)

const TestRecipeId = "2529831396465550" // Tasty pancakes
const TestUserId = "bcfce3497b42e48f1210823471c1312f"
const TestSearchQuery = "pizza"

var c types.Client

func init() {
	c = client.NewClient()
}
