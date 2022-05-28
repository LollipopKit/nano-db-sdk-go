package nanodbsdkgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var (
	client = &http.Client{}
)

func (db *DB) httpDo(fn, path string, body []byte) (data []byte, err error) {
	req, err := http.NewRequest(fn, db.url+path, bytes.NewReader(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", db.cookie)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
