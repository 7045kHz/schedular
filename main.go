package main

import (
	"fmt"
//	"os"
"sync"
 
	"os/exec"
	"models/job"
	)

var wg sync.WaitGroup





func WgExec(j Job) {
 
	cmd := exec.Command(j.Script ,  j.Args ... ) 
	for _,v := range j.Env {
		cmd.Env = append(cmd.Env, v)
		fmt.Printf("WgExec %v %v Env = %v\n",j.Script, j.Args , v)
	}
	
 
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	err := cmd.Run() 
	if err != nil {
		fmt.Printf("CMD RUN Error: %v\n", err)
	}
	fmt.Printf("CMD RESULTS: %v\n",cmd.String())
	wg.Done()
}
func main()	{

	var JobCount int
	
	var Jobs []Job
	var Job Job
	
	Job.Script="CMD.EXE"
	Job.Args=append(Job.Args,"TEST2.BAT")
	Job.Args=append(Job.Args,"/x=1")
	Job.Env=append(Job.Env,"MY_VAR=J1")
	Jobs = append(Jobs,Job)

	Job.Env = nil
	Job.Args = nil
	Job.Script = ""
	Job.Script="TEST2.BAT"
	Job.Env=append(Job.Env,"MY_VAR=J2")
	Jobs = append(Jobs,Job)

	JobCount=len(Jobs)
	fmt.Println("start wait group")
	wg.Add(JobCount) // indicate we are going to wait for one thing
	for i,v := range Jobs {
		fmt.Printf("I = %v, V = %v\n",i,v)
		go WgExec(v)
	}
	fmt.Printf("Waiting for commands to be finished...\n")
	wg.Wait() // wait for all things to be done
	// end of program

	fmt.Printf("Commands finished...\n")
	
	

}