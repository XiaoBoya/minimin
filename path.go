package minimin

import (
	"os"
	"strings"
)

var (
	PATH = ""
)

// InitStorage 初始化存储路径
func InitStorage(path string) (err error) {
	if !PathExist(path) {
		err = os.Mkdir(path, SimpleDirPerm)
		if err == nil {
			PATH = path
		} else {
			return err
		}
	} else {
		PATH = path
	}
	return nil
}

// GetBasePath 获取当前的路径
func GetBasePath() string {
	switch PATH {
	case "":
		dir, _ := os.Getwd()
		return dir
	default:
		return PATH
	}
}

// PathJoin 路径拼接
func PathJoin(path string, filename string) string {
	filename = strings.Trim(filename, "/")
	path = strings.TrimRight(path, "/")
	return path + "/" + filename
}

// PathExist 路径存在
func PathExist(path string) (res bool) {
	var err error
	_, err = os.Stat(path)
	switch err {
	case nil:
		return true
	default:
		return false
	}
}
