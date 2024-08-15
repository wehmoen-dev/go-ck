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
    "github.com/wehmoen/go-ck"
)

const MyFavouriteRecipeId = "2529831396465550"

func main() {
    ck := client.NewClient()
    recipes, err := ck.GetRecipe(MyFavouriteRecipeId)
    if err != nil {
        fmt.Println(err)
    }
    for _, recipe := range recipes {
        fmt.Println(recipe.Title)
    }
}
```