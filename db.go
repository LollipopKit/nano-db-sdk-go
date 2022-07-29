package nanodbsdkgo

import (
	"errors"
	"fmt"

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

func (db *DB) Dirs(dbName string) ([]string, error) {
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

	colsStr, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var cols []string
	err = json.Unmarshal(colsStr, &cols)
	if err == nil {
		return cols, nil
	}
	return nil, errors.New(fmt.Sprintf("data type error: %v", resp.Data))
}

func (db *DB) Files(dbName, dir string) ([]string, error) {
	data, err := db.httpDo("GET", dbName+"/"+dir, nil)
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

	filesStr, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var files []string
	err = json.Unmarshal(filesStr, &files)
	if err == nil {
		return files, nil
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

func (db *DB) DeleteDir(dbName, dir string) error {
	data, err := db.httpDo("DELETE", dbName+"/"+dir, nil)
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
