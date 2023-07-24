package ndb

import (
	"path"
)

func (cl *client) Alive() bool {
	_, err := cl.Do("HEAD", "", nil)
	return err == nil
}

func (cl *client) Read(paths ...string) (data []byte, err error) {
	return cl.Do("GET", path.Join(paths...), nil)
}

func (cl *client) Write(db, dir, file string, data []byte) error {
	_, err := cl.Do("POST", path.Join(db, dir, file), data)
	return err
}

func (cl *client) Delete(paths ...string) error {
	_, err := cl.Do("DELETE", path.Join(paths...), nil)
	return err
}
