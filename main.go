package main

import (
	"fmt"
	//	"os"

	"sync"

	"github.com/7045kHz/schedular/controllers"
	"github.com/7045kHz/schedular/models"
	"github.com/7045kHz/schedular/utils"
)

var wg sync.WaitGroup

func main() {

	var JobCount int

	var Jobs []models.Job
	var Job models.Job
	// Test Job/Jobs only
	Job.Env = nil
	Job.Args = nil
	Job.Engine = "CMD.EXE"
	Job.Exec, _ = utils.GetPath("TEST1.BAT")
	Job.Args = append(Job.Args, "/x=1")
	Job.Env = append(Job.Env, "MY_VAR=J1")
	Jobs = append(Jobs, Job)

	Job.Env = nil
	Job.Args = nil
	Job.Exec, _ = utils.GetPath("TEST2.BAT")
	Job.Engine = "CMD.EXE"
	Job.Env = append(Job.Env, "MY_VAR=J2")
	Jobs = append(Jobs, Job)

	Job.Env = nil
	Job.Args = nil
	Job.Engine = "POWERSHELL.EXE"
	Job.Exec, _ = utils.GetPath("powershell.exe")
	args := []string{"-NoProfile", "-NonInteractive", ".\\TEST3.ps1"}
	Job.Args = append(Job.Args, args...)
	fmt.Printf("Job.Args from Powershell: %v\n", Job.Args)
	Jobs = append(Jobs, Job)

	JobCount = len(Jobs)
	// exe(Job.Exec, Job.Args)
	fmt.Printf("start wait group for [%d] jobs\n", JobCount)
	wg.Add(JobCount) // indicate we are going to wait for one thing
	for i, v := range Jobs {
		fmt.Printf("I = %v, V = %v\n", i, v)
		go controllers.WgExec(v, &wg)
	}
	fmt.Printf("Waiting for commands to be finished...\n")
	wg.Wait() // wait for all things to be done
	// end of program

	fmt.Printf("Commands finished...\n")

}
