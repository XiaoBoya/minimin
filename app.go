package minimin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// App app
type App struct {
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	Name    string `json:"name" yaml:"name"`
}

// AppList 每个工程目录下的appList
type AppList map[string]AppInfo

// AppInfo app 信息
type AppInfo struct {
	Name string `json:"name"`
}

func (a *App) getProjectPath() (path string, err error) {
	var localPath = GetBasePath()
	path = PathJoin(localPath, a.Project)
	_, err = os.Stat(path)
	return
}

// CancelApp 注销app
func (a *App) CancelApp(path string) (err error) {
	path = PathJoin(path, InfoDir+"/"+AppListFile)
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	var al AppList
	if err = json.Unmarshal(content, &al); err != nil {
		return
	}
	if _, ok := al[a.Name]; !ok {
		return errors.New("the app is not existed")
	}
	delete(al, a.Name)
	content, err = json.Marshal(al)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, content, SimpleDirPerm)
	return
}

// RegisterApp 注册app
func (a *App) RegisterApp(path string) (err error) {
	path = PathJoin(path, InfoDir+"/"+AppListFile)
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	var al AppList
	if err = json.Unmarshal(content, &al); err != nil {
		return
	}
	if _, ok := al[a.Name]; ok {
		return errors.New("the app is existed")
	}
	al[a.Name] = AppInfo{Name: a.Name}
	content, err = json.Marshal(al)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, content, SimpleDirPerm)
	return
}

// New 新建app
func (a *App) New() (err error) {
	var path string
	path, err = a.getProjectPath()
	if err != nil {
		return
	}
	var appPath = PathJoin(path, a.Name)
	if PathExist(appPath) {
		return os.ErrExist
	}
	if err = os.Mkdir(appPath, SimpleDirPerm); err != nil {
		return
	}
	var appInfoPath = PathJoin(appPath, InfoDir)
	if err = os.Mkdir(appInfoPath, SimpleDirPerm); err != nil {
		return
	}
	if err = a.RegisterApp(path); err != nil {
		_ = os.Remove(appPath)
		return
	}
	return nil
}

// Delete 删除app
func (a *App) Delete() (err error) {
	var path string
	path, err = a.getProjectPath()
	if err != nil {
		return
	}
	var appPath = PathJoin(path, a.Name)
	if !PathExist(appPath) {
		return os.ErrNotExist
	}
	if err = a.CancelApp(path); err != nil {
		return
	}
	if err = os.Remove(appPath); err != nil {
		_ = a.RegisterApp(path)
		return
	}
	return nil
}
