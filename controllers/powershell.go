package controllers

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"

	"github.com/7045kHz/schedular/models"
	"github.com/7045kHz/schedular/utils"
)

func GetPowerShellPath() (ps string, err error) {
	ps, err = exec.LookPath("powershell.exe")
	if err != nil {
		fmt.Printf("GetPowerShellPath() error: %v\n", err)
		return "", err
	}
	return ps, nil
}

func WgPsExec(j models.Job, wg *sync.WaitGroup) (stdOut string, stdErr string, err error) {

	// func (p *models.PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args := append([]string{"-NoProfile", "-NonInteractive"}, j.Args...)
	Path, _ := utils.GetPath("powershell.exe")
	cmd := exec.Command(Path, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	wg.Done()
	return
}
