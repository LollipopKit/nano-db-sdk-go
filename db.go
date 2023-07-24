package ndb

import (
	"path"

	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (cl *client) Alive() bool {
	_, err := cl.httpDo("HEAD", "", nil)
	return err == nil
}

func (cl *client) Read(db, dir, file string) (data []byte, err error) {
	return cl.httpDo("GET", path.Join(db, dir, file), nil)
}

func (cl *client) Write(db, dir, file string, data []byte) error {
	_, err := cl.httpDo("POST", path.Join(db, dir, file), data)
	return err
}

func (cl *client) Delete(db, dir, file string) error {
	_, err := cl.httpDo("DELETE", path.Join(db, dir, file), nil)
	return err
}

func (cl *client) Dirs(db string) ([]string, error) {
	data, err := cl.httpDo("GET", db, nil)
	if err != nil {
		return nil, err
	}

	var strs []string
	return strs, json.Unmarshal(data, &strs)
}

func (cl *client) Files(db, dir string) ([]string, error) {
	data, err := cl.httpDo("GET", path.Join(db, dir), nil)
	if err != nil {
		return nil, err
	}

	var strs []string
	return strs, json.Unmarshal(data, &strs)
}

func (cl *client) DeleteDB(db string) error {
	_, err := cl.httpDo("DELETE", db, nil)
	return err
}

func (cl *client) DeleteDir(db, dir string) error {
	_, err := cl.httpDo("DELETE", path.Join(db, dir), nil)
	return err
}
