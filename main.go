package main

import (
	"encoding/json"
	"fmt"

	//	"os"

	"sync"

	"github.com/7045kHz/schedular/controllers"
	"github.com/7045kHz/schedular/drivers/mssql"
	"github.com/7045kHz/schedular/models"
	"github.com/7045kHz/schedular/utils"
	"github.com/subosito/gotenv"
)

//var db *sql.DB

var wg sync.WaitGroup

func init() {
	gotenv.Load()

}
func main() {

	var JobCount int

	// Connect to DB
	mssqldb := mssql.ConnectDB()
	// Get All Jobs - for now
	DBJobs := mssql.GetAllJobs(mssqldb)
	var Jobs []models.Job

	// look over returned Jobs from DB and use Jobs.Job_Definition to create a
	// slice of Job(s)
	for _, k := range DBJobs {
		var Job models.Job
		json.Unmarshal([]byte(k.Job_Definition), &Job)

		Job.Exec, _ = utils.GetPath(Job.Exec)
		TmpJob, _ := json.MarshalIndent(&Job, "", "	")
		if Job.Verbose == 1 {

			fmt.Printf("Job_Definition JSON: %v\n", string(TmpJob))
		}
		//fmt.Printf("Appending: %v\n", &Job.Args)
		Jobs = append(Jobs, Job)

	}
	//	fmt.Printf("Jobs after append: %v\n", Jobs)
	JobCount = len(Jobs)
	fmt.Printf("Starting wait group for [%d] jobs\n", JobCount)

	wg.Add(JobCount) // indicate we are going to wait for one thing
	for i, v := range Jobs {
		if v.Verbose == 1 {
			fmt.Printf("I = %v, V = %v\n", i, v)
		}
		go controllers.WgExec(v, &wg)
	}

	wg.Wait() // wait for all things to be done
	// end of program

	fmt.Printf("Wait Group Commands finished...\n")

}
