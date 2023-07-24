package ndb

import "strings"

type client struct {
	url   string
	token string
}

func NewClient(url, cookie string) *client {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return &client{
		url:   url,
		token: cookie,
	}
}
