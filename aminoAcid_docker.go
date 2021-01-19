package minimin

import (
	"log"
	"os/exec"
	"strings"
)

type DockerBuildParams struct {
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Cell     string `json:"-"`
	Repo     string `json:"repo"`
	User     string `json:"user"`
	Password string `json:"password"`
	HubAddr  string `json:"hub_addr,omitempty"`
}

func DockerBuild(obj DockerBuildParams) (err error) {
	var theCmd = exec.Command("docker", "build", "-t", obj.Name+":"+obj.Tag, obj.Cell+"/"+obj.Repo)
	var theCmdList = []string{"docker", "build", "-t", obj.Name + ":" + obj.Tag, obj.Cell + "/" + obj.Repo}
	log.Println(strings.Join(theCmdList, " "))
	if err = ExecCmd(theCmd); err != nil {
		return err
	}
	var loginCmd *exec.Cmd
	if obj.HubAddr != "" {
		loginCmd = exec.Command("docker", "login", "-u", obj.User, "--password-stdin")
	} else {
		loginCmd = exec.Command("docker", "login", obj.HubAddr, "-u", obj.User, "--password-stdin")
	}
	loginCmd.Stdin = strings.NewReader(obj.Password)
	if err = ExecCmd(loginCmd); err != nil {
		log.Fatalln("login\n" + err.Error())
		return err
	}
	return nil
}
