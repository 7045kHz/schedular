package main

import (
	"fmt"
	//	"os"

	"os/exec"
	"sync"

	"github.com/7045kHz/schedular/controllers"
	"github.com/7045kHz/schedular/models"
)

var wg sync.WaitGroup

func main() {

	var JobCount int

	var Jobs []models.Job
	var Job models.Job

	Job.Engine = "CMD.EXE"
	Job.Script = "TEST2.BAT"
	Job.Args = append(Job.Args, "/x=1")
	Job.Env = append(Job.Env, "MY_VAR=J1")
	Jobs = append(Jobs, Job)

	Job.Env = nil
	Job.Args = nil
	Job.Script = ""
	Job.Engine = ""
	Job.Script = "TEST2.BAT"
	Job.Env = append(Job.Env, "MY_VAR=J2")
	Jobs = append(Jobs, Job)

	Job.Env = nil
	Job.Args = nil
	Job.Script = ""
	Job.Engine = ""
	Job.Engine, _ = exec.LookPath("powershell.exe")

	Job.Script = "TEST2.BAT"

	Job.Args = append(Job.Args, "powershell.exe")

	Jobs = append(Jobs, Job)

	JobCount = len(Jobs)
	fmt.Println("start wait group")
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
