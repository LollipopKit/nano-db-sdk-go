package nanodbsdkgo

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (db *DB) Alive() bool {
	_, err := db.httpDo("HEAD", "", nil)
	return err == nil
}

func (db *DB) Status() (status string, err error) {
	data, err := db.httpDo("GET", "", nil)
	if err != nil {
		return
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}

	status = resp.Data.(string)
	if resp.Code != 200 {
		err = errors.New(status)
		return
	}

	return
}

func (db *DB) Read(path string, mod any) (err error) {
	data, err := db.httpDo("GET", path, nil)
	if err != nil {
		return
	}

	var resp Resp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
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

func (db *DB) Write(path string, mod any) error {
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

	// 这里不能直接用resp.Data.([]string)来转换，因为resp.Data是interface{}类型
	// 下方Files()方法也是如此
	// -------------
	// `go test`表明
	// 此处用json转换比`resp.Data.([]interface{})`更快
	dirsStr, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var dirs []string
	err = json.Unmarshal(dirsStr, &dirs)
	if err == nil {
		return dirs, nil
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

// 搜索某个dir内，所有[gjson.Get(_,p).Exists() == true]的FILE。
// 如果正则不为空，仅返回正则匹配成功的FILE。
func (db *DB) Search(dbName, dir, gjsonPath, valueRegex string) ([]any, error) {
	p := fmt.Sprintf("%s/%s?path=%s&value=%s", dbName, dir, gjsonPath, valueRegex)
	data, err := db.httpDo("POST", p, nil)
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

	dataStr, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var datas []any
	err = json.Unmarshal(dataStr, &datas)
	if err == nil {
		return datas, nil
	}
	return nil, errors.New("data type error: " + resp.Data.(string))
}
