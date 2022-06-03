package nanodbsdkgo

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

const (
	errNoDocStr = "db.Read(): no document"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	ErrNoDoc = errors.New("db.Read(): no document")
)

func (db *DB) Alive() bool {
	_, err := db.httpDo("HEAD", "", nil)
	return err == nil
}

func (db *DB) Status() (string, error) {
	data, err := db.httpDo("GET", "", nil)
	if err != nil {
		return "", err
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return "", err
	}

	body := resp.Data.(string)
	if resp.Code != 200 {
		return "", errors.New(body)
	}

	return body, nil
}
	

func (db *DB) Read(path string, mod interface{}) error {
	data, err := db.httpDo("GET", path, nil)
	if err != nil {
		return err
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		errStr := resp.Data.(string)
		if errStr == errNoDocStr {
			return ErrNoDoc
		}
		return errors.New(errStr)
	}

	data, err = json.Marshal(resp.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, mod)
}

func (db *DB) Write(path string, mod interface{}) error {
	data, err := json.Marshal(mod)
	if err != nil {
		return err
	}
	data, err = db.httpDo("POST", path, data)
	if err != nil {
		return err
	}
	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Data.(string))
	}
	return nil
}

func (db *DB) Delete(path string) error {
	data, err := db.httpDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Data.(string))
	}
	return nil
}

func (db *DB) Cols(dbName string) ([]string, error) {
	data, err := db.httpDo("GET", dbName, nil)
	if err != nil {
		return nil, err
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, errors.New(resp.Data.(string))
	}

	cols, ok := resp.Data.([]string)
	if ok {
		return cols, nil
	}
	return nil, errors.New("data type error: " + resp.Data.(string))
}

func (db *DB) Files(dbName, col string) ([]string, error) {
	data, err := db.httpDo("GET", dbName+"/"+col, nil)
	if err != nil {
		return nil, err
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, errors.New(resp.Data.(string))
	}

	ids, ok := resp.Data.([]string)
	if ok {
		return ids, nil
	}
	return nil, errors.New("data type error: " + resp.Data.(string))
}

func (db *DB) DeleteDB(dbName string) error {
	data, err := db.httpDo("DELETE", dbName, nil)
	if err != nil {
		return err
	}
	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Data.(string))
	}
	return nil
}

func (db *DB) DeleteCol(dbName, col string) error {
	data, err := db.httpDo("DELETE", dbName+"/"+col, nil)
	if err != nil {
		return err
	}
	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Data.(string))
	}
	return nil
}
