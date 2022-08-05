package controllers

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"

	"github.com/7045kHz/schedular/models"
)

func WgExec(j models.Job, wg *sync.WaitGroup) (stdOut string, stdErr string, err error) {

	cmd := exec.Command(j.Exec, j.Args...)
	for _, v := range j.Env {
		cmd.Env = append(cmd.Env, v)
		fmt.Printf("WgExec %v %v Env = %v\n", j.Exec, j.Args, v)
	}

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("CMD RUN Error: %v\n", err)
	}
	fmt.Printf("CMD RESULTS: %v\n", cmd.String())
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	stdOut, stdErr = stdout.String(), stderr.String()
	wg.Done()
	return

}
