package nanodbsdkgo

import "strings"

type DB struct {
	url    string
	cookie string
}

func NewDB(url, cookie string) *DB {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return &DB{
		url:    url,
		cookie: cookie,
	}
}

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
