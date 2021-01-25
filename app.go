package minimin

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func (a *App) appInfoPath() (path string) {
	path = filepath.Join(a.Path, InfoDir)
	return
}

func (a *App) appConfigYamlPath() (path string) {
	path = filepath.Join(a.Path, InfoDir, ConfigYamlFile)
	return
}

func (a *App) appConfigJsonPath() (path string) {
	path = filepath.Join(a.Path, InfoDir, ConfigJsonFile)
	return
}

func (a *App) appRunLogJsonPath() (path string) {
	path = filepath.Join(a.Path, InfoDir, RunLogJsonFile)
	return
}

func (a *App) getProjectPath() (path string, err error) {
	var localPath = GetBasePath()
	path = filepath.Join(localPath, a.Project.Name)
	_, err = os.Stat(path)
	return
}

// CancelApp 注销app
func (a *App) CancelApp(path string) (err error) {
	path = filepath.Join(path, InfoDir+"/"+AppListFile)
	var content []byte
	var littleLock = sync.Mutex{}
	littleLock.Lock()
	defer littleLock.Unlock()
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
	path = filepath.Join(path, InfoDir+"/"+AppListFile)
	var content []byte
	var littleLock = sync.Mutex{}
	littleLock.Lock()
	defer littleLock.Unlock()
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
	a.Path = filepath.Join(path, a.Name)
	if PathExist(a.Path) {
		return os.ErrExist
	}
	if err = os.Mkdir(a.Path, SimpleDirPerm); err != nil {
		return
	}
	var appInfoPath = a.appInfoPath()
	if err = os.Mkdir(appInfoPath, SimpleDirPerm); err != nil {
		return
	}
	if err = a.RegisterApp(path); err != nil {
		_ = os.Remove(a.Path)
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
	if !PathExist(a.Path) {
		return os.ErrNotExist
	}
	if err = a.CancelApp(path); err != nil {
		return
	}
	if err = os.Remove(a.Path); err != nil {
		_ = a.RegisterApp(path)
		return
	}
	return nil
}

// LoadConfigYamlByFile 通过文件加载配置（CI/CD运行使用的文件）
func (a *App) LoadConfigYamlByFile(path string) (err error) {
	if _, err = YamlLoad(path); err != nil {
		return
	}
	var content []byte
	if content, err = ioutil.ReadFile(path); err != nil {
		return
	}
	err = ioutil.WriteFile(a.appConfigYamlPath(), content, SimpleFilePerm)
	return
}

// LoadConfigYamlByByte 通过内容加载配置（CI/CD运行使用的文件）
func (a *App) LoadConfigYamlByByte(content []byte) (err error) {
	if _, err = YamlLoadByByte(content); err != nil {
		return
	}
	err = ioutil.WriteFile(a.appConfigYamlPath(), content, SimpleFilePerm)
	return
}

// GetConfigYaml 获取运行配置
func (a *App) GetConfigYaml() (mf *MinFile, err error) {
	mf, err = YamlLoad(a.appConfigYamlPath())
	return
}

// LoadConfigJsonByFile 通过文件加载配置（管理的json文件）
func (a *App) LoadConfigJsonByFile(path string) (err error) {
	var aac AppAdminConfig
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &aac); err != nil {
		return
	}
	err = ioutil.WriteFile(a.appConfigJsonPath(), content, SimpleFilePerm)
	return
}

// LoadConfigJsonByFile 通过文件加载配置（管理的json文件）
func (a *App) LoadConfigJsonByByte(content []byte) (err error) {
	var aac AppAdminConfig
	if err = json.Unmarshal(content, &aac); err != nil {
		return
	}
	err = ioutil.WriteFile(a.appConfigJsonPath(), content, SimpleFilePerm)
	return
}

// GetConfigJson 获取配置
func (a *App) GetConfigJson() (aac *AppAdminConfig, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(a.appConfigJsonPath()); err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &aac)
	return
}

// Run 运行app
func (a *App) Run() (plObj *DNA, err error) {
	var l = sync.Mutex{}
	var num int
	l.Lock()
	defer l.Unlock()
	var content []byte
	content, err = ioutil.ReadFile(a.appRunLogJsonPath())
	if err != nil {
		return
	}
	var rljList []RunLogJson
	if err = json.Unmarshal(content, &rljList); err != nil {
		return
	}
	num = len(rljList) + 1
	plObj, err = a.NewProductionLine(num)
	if err != nil {
		return
	}
	rljList = append(rljList, RunLogJson{
		Num:       num,
		StartTime: sql.NullTime{Time: time.Now(), Valid: true},
		Status:    Queue,
	})
	b, _ := json.Marshal(rljList)
	err = ioutil.WriteFile(a.appRunLogJsonPath(), b, SimpleFilePerm)
	if err != nil {
		os.Remove(filepath.Join(a.Path, strconv.Itoa(num)))
	}
	return
}

// NewProductionLine 新启动一条生产线
func (a *App) NewProductionLine(num int) (pl *DNA, err error) {
	var numStr = strconv.Itoa(num)
	var thePLPath = filepath.Join(a.Path, numStr)
	if err = os.Mkdir(thePLPath, SimpleDirPerm); err != nil {
		return nil, err
	}
	var content []byte
	content, err = ioutil.ReadFile(filepath.Join(a.Project.projectInfoPath(), ConfigYamlFile))
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(filepath.Join(thePLPath, ConfigYamlFile), content, SimpleFilePerm)
	if err != nil {
		os.Remove(thePLPath)
	}
	pl = &DNA{Num: numStr, Path: thePLPath}
	return
}

// RunLogJson 运行记录的json记录文件
type RunLogJson struct {
	Num       int          `json:"num"`
	StartTime sql.NullTime `json:"start_time"`
	EndTime   sql.NullTime `json:"end_time"`
	Status    Status       `json:"status"`
}
