package minimin

import (
	"io/ioutil"
	"os"
)

type Project struct {
	Name string `json:"name" yaml:"name"`
}

func (p *Project) New() (err error) {
	var localPath = GetBasePath()
	var path = PathJoin(localPath, p.Name)
	if PathExist(path) {
		return os.ErrExist
	}
	if err = os.Mkdir(path, SimpleFilePerm); err != nil {
		return
	}
	path = PathJoin(path, InfoDir)
	if err = os.Mkdir(path, SimpleFilePerm); err != nil {
		return
	}
	var appConfPath = PathJoin(path, AppListFile)
	err = ioutil.WriteFile(appConfPath, nil, 0666)
	return nil
}

func (p *Project) Delete() (err error) {
	var localPath = GetBasePath()
	var path = PathJoin(localPath, p.Name)
	if PathNotExist(path) {
		return os.ErrNotExist
	}
	err = os.Remove(path)
	return
}
