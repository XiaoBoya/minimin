package minimin

import (
	"encoding/json"
	"log"
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
