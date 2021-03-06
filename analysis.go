package minimin

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// YamlLoad 读取配置文件（yaml）类型
func YamlLoad(path string) (anriFile *MinFile, err error) {
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &anriFile)
	return
}

// YamlLoadByByte 读取配置文件
func YamlLoadByByte(content []byte) (anriFile *MinFile, err error) {
	err = yaml.Unmarshal(content, &anriFile)
	return
}

// HandleMPs 处理处理单元
func HandleMPs(MPList []MP) (MPB []MPForWork) {
	var multi *MPForWork
	for _, MPObj := range MPList {
		if MPObj.Chain == "" {
			if multi != nil {
				MPB = append(MPB, *multi)
			}
			temporaryMP := MPObj
			MPB = append(MPB, MPForWork{
				Single: &temporaryMP,
				Multi:  nil,
			})
			multi = nil
		} else {
			if multi == nil {
				multi = &MPForWork{
					nil,
					map[string][]MP{},
				}
			}
			if _, ok := multi.Multi[MPObj.Chain]; !ok {
				multi.Multi[MPObj.Chain] = []MP{MPObj}
			} else {
				multi.Multi[MPObj.Chain] = append(multi.Multi[MPObj.Chain], MPObj)
			}
		}
	}
	if multi != nil {
		MPB = append(MPB, *multi)
	}
	return
}
