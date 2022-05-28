package nanodbsdkgo

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

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
		return errors.New(resp.Data.(string))
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

func (db *DB) Exist(path string) (bool, error) {
	data, err := db.httpDo("HEAD", path, nil)
	if err != nil {
		return false, err
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return false, err
	}
	if resp.Code == 200 {
		exist, ok := resp.Data.(bool)
		if ok {
			return exist, nil
		}
		return false, errors.New("data type error")
	}
	return false, errors.New(resp.Data.(string))
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

func (db *DB) Ids(dbName, col string) ([]string, error) {
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
