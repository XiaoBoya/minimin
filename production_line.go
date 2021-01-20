package minimin

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func (d *DNA) LoadMinFile() (minFile *MinFile, err error) {
	var content []byte
	content, err = ioutil.ReadFile(PathJoin(d.Path, EnvJsonFile))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &minFile)
	return
}

// HandleMinFile 配置文件转换到运行的结构体
func (d *DNA) HandleMinFile(file MinFile) (err error) {
	d.Env = file.Env
	var envB []byte
	if envB, err = json.Marshal(file.Env); err != nil {
		return
	}
	if err = ioutil.WriteFile(PathJoin(d.Path, EnvJsonFile), envB, SimpleFilePerm); err != nil {
		return
	}
	for _, geneObj := range file.Genes {
		var newGeneObj GeneForWork
		newGeneObj.Status = Queue
		newGeneObj.Name = geneObj.Name
		newGeneObj.Ignore = geneObj.Ignore
		newGeneObj.When = geneObj.When
		newGeneObj.MPs = HandleMPs(geneObj.MPs)
		d.Genes = append(d.Genes, newGeneObj)
	}
	return nil
}

func (d *DNA) work() (err error) {
	for _, geneObj := range d.Genes {
		for _, mpObj := range geneObj.MPs {
			switch mpObj.Multi {
			case nil:
				if _, err = mpObj.Single.Synthesis(d.Path); err != nil {
					if mpObj.Single.Ignore {
						log.Println(err.Error())
					} else {
						log.Fatalln(err.Error())
						return
					}
				}
			default:
				for _, mpChildList := range mpObj.Multi {
					go ListSynthesis(mpChildList, d.Path)
				}
			}
		}
	}
	return
}

func (d *DNA) Work() (err error) {
	var minFile *MinFile
	if minFile, err = d.LoadMinFile(); err != nil {
		return
	}
	if err = d.HandleMinFile(*minFile); err != nil {
		return
	}
	err = d.Work()
	return
}
