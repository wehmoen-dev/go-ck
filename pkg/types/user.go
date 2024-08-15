package types

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

type User struct {
	restClient *resty.Client

	Id                     string   `json:"id"`
	Username               string   `json:"username"`
	Rank                   int      `json:"rank"`
	Role                   string   `json:"role"`
	HasAvatar              bool     `json:"hasAvatar"`
	HasPaid                bool     `json:"hasPaid"`
	Deleted                bool     `json:"deleted"`
	AvatarImageUrlTemplate string   `json:"avatarImageUrlTemplate,omitempty"`
	AvatarImageUrl         string   `json:"avatarImageUrl,omitempty"`
	IsVerified             bool     `json:"isVerified"`
	DisplayName            string   `json:"displayName"`
	Profile                *Profile `json:"profile,omitempty"`
}

func (u *User) Parse() {
	u.AvatarImageUrl = strings.ReplaceAll(u.AvatarImageUrlTemplate, "<format>", UserAvatarImageFormat)
	u.AvatarImageUrlTemplate = ""
}

func (u *User) SetRestClient(client *resty.Client) {
	u.restClient = client
}

func (u *User) GetProfile() error {
	if u.Profile != nil {
		return nil
	}

	var result *Profile

	response, err := u.restClient.R().SetResult(&result).Get(fmt.Sprintf("/users/%s/profile", u.Id))

	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		return fmt.Errorf("received non 2xx response: %d - %v", response.StatusCode(), response)
	}

	u.Profile = result
	return nil
}

type Profile struct {
	CookKnowHow     int    `json:"cookKnowHow"`
	Username        string `json:"username"`
	MemberSince     string `json:"memberSince"`
	FormOfNutrition string `json:"formOfNutrition"`
	HasAvatar       bool   `json:"hasAvatar"`
	Audience        string `json:"audience"`
	IsComUser       bool   `json:"isComUser"`
}
