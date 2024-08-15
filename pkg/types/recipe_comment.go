package types

import "time"

type Comment struct {
	Id           string      `json:"id"`
	Owner        *User       `json:"owner"`
	CreatedAt    time.Time   `json:"createdAt"`
	Text         string      `json:"text"`
	Helpful      bool        `json:"helpful"`
	HelpfulCount int         `json:"helpfulCount"`
	ParentId     *string     `json:"parentId"`
	Children     interface{} `json:"children"`
}
