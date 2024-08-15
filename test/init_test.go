package test

import (
	"github.com/wehmoen/go-ck/pkg/client"
	"github.com/wehmoen/go-ck/pkg/types"
)

var c types.Client

func init() {
	c = client.NewClient()
}
