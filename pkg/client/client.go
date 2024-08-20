package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/wehmoen-dev/go-ck/pkg/types"
)

const ApiBaseUrl = "https://api.chefkoch.de/v2"

type c struct {
	restClient *resty.Client
}

func NewClient() types.Client {

	restClient := resty.New()
	restClient.SetBaseURL(ApiBaseUrl)

	return &c{
		restClient: restClient,
	}
}
