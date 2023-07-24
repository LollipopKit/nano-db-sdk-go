package ndb

import (
	"fmt"

	"github.com/lollipopkit/gommon/http"
)

func (db *client) httpDo(fn, path string, body []byte) (data []byte, err error) {
	resp, code, err := http.Do(fn, db.url+path, body, map[string]string{
		"NanoDB": db.token,
	})
	if code != 200 {
		err = fmt.Errorf("code: %d, resp: %s", code, string(resp))
	}
	return resp, err
}
