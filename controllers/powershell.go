package controllers

import (
	"bytes"
	"os/exec"
)

func PowerShellNew() *modules.PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &modules.PowerShell{
		powerShell: ps,
	}
}

func (p *modules.PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}
