package client

import (
	"fmt"
	"github.com/wehmoen-dev/go-ck/pkg/types"
)

func (c *c) GetUser(userId string, includeProfile bool) (*types.User, error) {
	var user *types.User

	response, err := c.restClient.R().SetResult(&user).Get(fmt.Sprintf("/users/%s", userId))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("received non 2xx response: %d - %v", response.StatusCode(), response)
	}

	user.SetRestClient(c.restClient)
	user.Parse()

	if includeProfile {
		err = user.GetProfile()
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
