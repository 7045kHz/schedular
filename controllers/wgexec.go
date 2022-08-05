package controllers

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/7045kHz/schedular/models"
)

func WgExec(j models.Job, wg *sync.WaitGroup) {

	cmd := exec.Command(j.Script, j.Args...)
	for _, v := range j.Env {
		cmd.Env = append(cmd.Env, v)
		fmt.Printf("WgExec %v %v Env = %v\n", j.Script, j.Args, v)
	}

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("CMD RUN Error: %v\n", err)
	}
	fmt.Printf("CMD RESULTS: %v\n", cmd.String())
	wg.Done()
	return
}
