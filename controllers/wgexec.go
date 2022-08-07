package controllers

import (
	//	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/7045kHz/schedular/models"
)

type Controller struct{}

func WgExec(j models.Job, wg *sync.WaitGroup) (stdOut string, stdErr string, err error) {

	cmd := exec.Command(j.Exec)
	for _, v := range j.Env {

		cmd.Env = append(cmd.Env, v)
		if j.Verbose == 1 {
			fmt.Printf("Adding Env for Job [%v], Env = %v\n", j.Name, v)

		}
	}

	for _, v := range j.Args {
		cmd.Args = append(cmd.Args, v)
		if j.Verbose == 1 {
			fmt.Printf("Adding Args for Job [%v],  Args - %v\n", j.Name, v)
		}
	}
	if j.Verbose == 1 {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	if err != nil {
		fmt.Printf("JOB [%s] RUN Error: %v\n", j.Name, err)
	}
	fmt.Printf("JOB [%s] EXECUTED: %v %v \n", j.Name, j.Exec, j.Args)

	wg.Done()
	return
}
