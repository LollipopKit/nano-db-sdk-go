package nanodbsdkgo

type DB struct {
	url    string
	cookie string
}

func NewDB(url, cookie string) *DB {
	return &DB{
		url:    url,
		cookie: cookie,
	}
}

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
