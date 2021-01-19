package minimin

import (
	"errors"
	"os/exec"
	"strings"
)

type ShellParams struct {
	Cmd string `json:"cmd"`
}

func Shell(obj ShellParams) (err error) {
	var shellList = strings.Split(obj.Cmd, " ")
	var theCmd *exec.Cmd
	switch len(shellList) {
	case 0:
		return errors.New("no cmd")
	case 1:
		theCmd = exec.Command(obj.Cmd)
	default:
		theCmd = exec.Command(shellList[0], shellList[1:]...)
	}
	if err = ExecCmd(theCmd); err != nil {
		return
	}
	return err
}
