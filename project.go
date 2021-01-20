package minimin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Project 工程
type Project struct {
	Name string `json:"name" yaml:"name"`
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}

func (p *Project) projectInfoPath() (path string) {
	path = PathJoin(p.Path, InfoDir)
	return
}

func (p *Project) GetPath() {
	if p.Path == "" {
		var localPath = GetBasePath()
		p.Path = PathJoin(localPath, p.Name)
	}
}

// New 新建工程
func (p *Project) New() (err error) {
	p.GetPath()
	if PathExist(p.Path) {
		return os.ErrExist
	}
	if err = os.Mkdir(p.Path, SimpleDirPerm); err != nil {
		return
	}
	var path = p.projectInfoPath()
	if err = os.Mkdir(path, SimpleDirPerm); err != nil {
		return
	}
	var appConfPath = PathJoin(path, AppListFile)
	err = ioutil.WriteFile(appConfPath, []byte("{}"), SimpleFilePerm)
	return nil
}

// Delete 删除工程
func (p *Project) Delete() (err error) {
	p.GetPath()
	if !PathExist(p.Path) {
		return os.ErrNotExist
	}
	err = os.Remove(p.Path)
	return
}

// NewApp 新建app
func (p *Project) NewApp(name string) (a *App, err error) {
	a = &App{
		Project: *p,
		Name:    name,
	}
	if err = a.New(); err != nil {
		return nil, err
	}
	return
}

// GetApp 获取app
func (p *Project) GetApp(name string) (a *App, err error) {
	p.GetPath()
	var path = p.projectInfoPath()
	var infoPath = PathJoin(path, InfoDir, AppListFile)
	var content []byte
	content, err = ioutil.ReadFile(infoPath)
	if err != nil {
		return nil, err
	}
	var al AppList
	if err = json.Unmarshal(content, &al); err != nil {
		return nil, err
	}
	if _, ok := al[name]; !ok {
		return nil, errors.New("the app not existed")
	}
	a = &App{
		Project: *p,
		Name:    name,
	}
	a.Path = PathJoin(p.Path, a.Name)
	return
}
