# Go-CK

This package provides a simple rest client for the `chefkoch.de` API.

## Installation

```bash
go get -u github.com/wehmoen/go-ck
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/wehmoen/go-ck/pkg/client"
)

const MyFavouriteRecipeId = "2529831396465550"

func main() {
	ck := client.NewClient()
	recipe, err := ck.GetRecipe(MyFavouriteRecipeId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(recipe.Title)
}
```
