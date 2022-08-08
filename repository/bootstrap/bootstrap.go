package bootstrap

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/7045kHz/schedular/models"
	"github.com/7045kHz/schedular/repository/schedule"
	"github.com/7045kHz/schedular/utils"
)

func GetOptions(mssqldb *sql.DB) {
	var alist bool
	var help bool
	var debug bool

	flag.BoolVar(&alist, "l", false, "List all scheduled jobs")
	flag.BoolVar(&debug, "d", false, "Dumps all schedule information")
	flag.BoolVar(&help, "h", false, "Command Line Help")

	flag.Parse()
	if help {
		fmt.Println("Information: Running schedular without options will check the schedule and execute any jobs scheduled for this time.")
		fmt.Println("Usage: schedular -h")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if alist {
		jobs := schedule.GetAllJobs(mssqldb)

		// fmt.Printf("%+v\n", jobs)
		for _, k := range jobs {

			fmt.Printf("\nJob_ID=%d\nName=%s\nExecution_Server=%s\nCreated_On=%s\nEnabled=%d\nStarted=%s\nFinished=%s\nLast_Updated=%s\n\n", k.Job_Id, k.Name, k.Execution_Server, k.Created_On, k.Enabled, k.Started, k.Finished, k.Last_Updated)

		}

		os.Exit(1)
	}
	if debug {
		jobs := schedule.GetAllJobs(mssqldb)

		// fmt.Printf("%+v\n", jobs)
		for _, k := range jobs {
			var Job models.Job
			json.Unmarshal([]byte(k.Job_Definition), &Job)
			Job.Job_Id = k.Job_Id
			Job.Exec, _ = utils.GetPath(Job.Exec)
			TmpJob, _ := json.MarshalIndent(&Job, "", "	")

			fmt.Printf("\nJob_ID=%d\nName=%s\nExecution_Server=%s\nCreated_On=%s\nEnabled=%d\nStarted=%s\nFinished=%s\nLast_Updated=%s\nJob_Definition='\n%+v'\n\n", k.Job_Id, k.Name, k.Execution_Server, k.Created_On, k.Enabled, k.Started, k.Finished, k.Last_Updated, string(TmpJob))

		}

		os.Exit(1)
	}
}
