package minimin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (mp *MP) SetStatus(status Status) {
	mp.Status = status
}

func (mp *MP) Synthesis(cellSpace string) (status string, err error) {
	mp.SetStatus(Init)
	defer func() {
		if err != nil {
			mp.SetStatus(Success)
		} else {
			mp.SetStatus(Fail)
		}
	}()
	paramsByte, _ := json.Marshal(mp.Params)
	if strings.HasPrefix(mp.AminoAcid, "@") {
		return
	}
	switch mp.AminoAcid {
	case "shell":
		var sp ShellParams
		if err = json.Unmarshal(paramsByte, &sp); err != nil {
			return
		}
		mp.Status = Run
		err = Shell(sp)
	case "git":
		var gp GitParams
		if err = json.Unmarshal(paramsByte, &gp); err != nil {
			log.Fatalln(err.Error())
			return
		}
		gp.Cell = cellSpace
		mp.Status = Run
		err = GitPull(gp)
	case "docker":
		var dp DockerBuildParams
		if err = json.Unmarshal(paramsByte, &dp); err != nil {
			return
		}
		dp.Cell = cellSpace
		mp.Status = Run
		err = DockerBuild(dp)
	case "email":
		var ep MailboxConf
		if err = json.Unmarshal(paramsByte, &ep); err != nil {
			return
		}
		mp.Status = Run
		err = SendEmail(ep)
	}
	return
}

func ListSynthesis(list []MP, cellSpace string) (err error) {
	for _, mpObj := range list {
		if _, err = mpObj.Synthesis(cellSpace); err != nil {
			if mpObj.Ignore {
				log.Println(err.Error())
			} else {
				log.Fatalln(err.Error())
				return
			}
		}
	}
	return
}

func Work(content []byte) (err error) {
	var af *MinFile
	af, err = YamlLoadByByte(content)
	if err != nil {
		return
	}
	var dnaObj *DNA
	dnaObj, err = HandleMinFile(*af)
	if err != nil {
		return
	}

	// create cell space
	var cellSpace = "cell/" + af.Name
	err = os.Mkdir(cellSpace, 0777)
	if err != nil {
		return
	}

	// write env file
	var b []byte
	b, _ = json.Marshal(af.Env)
	err = ioutil.WriteFile(cellSpace+"/env.json", b, 0777)
	if err != nil {
		return
	}

	// translate genes
	for _, geneObj := range dnaObj.Genes {
		for _, mpObj := range geneObj.MPs {
			switch mpObj.Multi {
			case nil:
				if _, err = mpObj.Single.Synthesis(cellSpace); err != nil {
					if mpObj.Single.Ignore {
						log.Println(err.Error())
					} else {
						log.Fatalln(err.Error())
						return
					}
				}
			default:
				for _, mpChildList := range mpObj.Multi {
					go ListSynthesis(mpChildList, cellSpace)
				}
			}
		}
	}
	return
}
