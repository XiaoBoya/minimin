package minimin

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func handleStdout(stdout io.ReadCloser) {
	reader := bufio.NewReader(stdout)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil || err == io.EOF {
			return
		}
		log.Println(string(lineBytes))
	}
}

func handleStderr(stderr io.ReadCloser) (errInfo string) {
	reader := bufio.NewReader(stderr)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil || err == io.EOF {
			return
		}
		log.Fatalln(string(lineBytes))
	}
}

func ExecCmd(theCmd *exec.Cmd) (err error) {
	var stdout io.ReadCloser
	var stderr io.ReadCloser
	if stdout, err = theCmd.StdoutPipe(); err != nil {
		return err
	}
	if stderr, err = theCmd.StderrPipe(); err != nil {
		return err
	}
	if err = theCmd.Start(); err != nil {
		return
	}
	go handleStdout(stdout)
	go handleStderr(stderr)
	err = theCmd.Wait()
	return err
}
