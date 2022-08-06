package controllers

import (
	//	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/7045kHz/schedular/models"
)

// (stdOut string, stdErr string, err error)
func WgExec(j models.Job, wg *sync.WaitGroup) (stdOut string, stdErr string, err error) {

	cmd := exec.Command(j.Exec, j.Args...)
	for _, v := range j.Env {
		cmd.Env = append(cmd.Env, v)
		fmt.Printf("WgExec %v, Args - %v, Env = %v\n", j.Exec, j.Args, v)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// var stdout bytes.Buffer
	// var stderr bytes.Buffer
	//cmd.Stdout = &stdout
	//cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("CMD RUN Error: %v\n", err)
	}
	fmt.Printf("CMD EXECUTED: %v\n", cmd.String())

	// stdOut, stdErr = stdout.String(), stderr.String()

	// fmt.Printf("STDOUT: %s\n", stdOut)
	// fmt.Printf("STDERR: %s\n", stdErr)
	wg.Done()
	return
}
