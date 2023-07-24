package ndb

import (
	"fmt"
	"strings"

	"github.com/lollipopkit/gommon/http"
)

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

func (db *client) Do(fn, path string, body []byte) (data []byte, err error) {
	resp, code, err := http.Do(fn, db.url+path, body, map[string]string{
		"NanoDB": db.token,
	})
	if code != 200 {
		err = fmt.Errorf("code: %d, resp: %s", code, string(resp))
	}
	return resp, err
}
