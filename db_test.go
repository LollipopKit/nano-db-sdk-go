package nanodbsdkgo_test

import (
	"testing"

	nanodbsdkgo "git.lolli.tech/lollipopkit/nano-db-sdk-go"
)

const (
	dbUrl = "http://localhost:3777"
	dbCookie = "n=bm92ZWw=; s=5bde31dc803625bcd0098e6e3d6bd07734dc8"
	rwdFilePath = "novel/23/27145"
	noSuchFilePath = "novel/23/27146"
)

var (
	db = nanodbsdkgo.NewDB(dbUrl, dbCookie)
)

func TestWrite(t *testing.T) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	err := db.Write(rwdFilePath, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestRead(t *testing.T) {
	var data map[string]interface{}
	err := db.Read(rwdFilePath, &data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestDelete(t *testing.T) {
	err := db.Delete(rwdFilePath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestReadNoSuchFile(t *testing.T) {
	var data map[string]interface{}
	err := db.Read(noSuchFilePath, &data)
	if err == nil {
		t.Fatal("should be error")
	}
	t.Log(err)
}

func TestDeleteNoSuchFile(t *testing.T) {
	err := db.Delete(noSuchFilePath)
	if err == nil {
		t.Fatal("should be error")
	}
	t.Log(err)
}

func TestDirs(t *testing.T) {
	dirs, err := db.Dirs("novel")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirs)
}

func TestFiles(t *testing.T) {
	files, err := db.Files("novel", "23")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total %d files", len(files))
}