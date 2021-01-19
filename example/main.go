package main

import (
	"github.com/XiaoBoya/minimin"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	if err = minimin.InitStorage("/Users/raymond/Desktop/xby"); err != nil {
		logrus.Error(err.Error())
	}
	logrus.Info(minimin.PATH)
}
