package main

import (
	"github.com/XiaoBoya/minimin"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	if err = minimin.InitStorage("/Users/raymond/Desktop/xby"); err != nil {
		logrus.Error("init project error:%s", err.Error())
		return
	}
	firstProject := minimin.Project{Name: "v1"}
	if err = firstProject.New(); err != nil {
		logrus.Errorf("create project error:%s", err.Error())
		return
	}
	var firstApp *minimin.App
	firstApp, err = firstProject.NewApp("app01")
	if err != nil {
		logrus.Errorf("create app error:%s", err.Error())
		return
	}
	if err = firstApp.LoadConfigYamlByFile("./test.yaml"); err != nil {
		logrus.Errorf("Load config file which type is yaml error:%s", err.Error())
		return
	}
}
