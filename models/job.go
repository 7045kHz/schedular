package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/7045kHz/schedular/utils"
)

type Job struct {
	Name    string   `json:"name"`
	Job_Id  int      `json:"job_id"`
	Engine  string   `json:"engine"`
	Exec    string   `json:"exec"`
	Verbose int      `json:"verbose"`
	Args    []string `json:"args"`
	Env     []string `json:"env"`
}

func (j Job) StartJob(mssqldb *sql.DB, now string) {
	var inventory_table = os.Getenv("INVENTORY_TABLE")
	fmt.Printf("In startjob for JID = %d\n", j.Job_Id)
	sqlSelect := fmt.Sprintf("UPDATE  %s SET [LAST_UPDATED]=CURRENT_TIMESTAMP, [FINISHED]='', [STARTED]='%s' WHERE Id=%d", inventory_table, now, j.Job_Id)
	// fmt.Printf("StartJob sql = %s\n", sqlSelect)
	stmt, err := mssqldb.Prepare(sqlSelect)
	utils.LogFatal(err)

	defer stmt.Close()
	_, err = stmt.Exec()
	utils.LogFatal(err)

}

func (j Job) FinishJob(mssqldb *sql.DB) {

	finished := utils.Now()
	var inventory_table = os.Getenv("INVENTORY_TABLE")
	sqlSelect := fmt.Sprintf("UPDATE %s SET [LAST_UPDATED]=CURRENT_TIMESTAMP, [FINISHED]='%s' WHERE Id=%d", inventory_table, finished, j.Job_Id)
	stmt, err := mssqldb.Prepare(sqlSelect)
	utils.LogFatal(err)

	defer stmt.Close()
	_, err = stmt.Exec()
	utils.LogFatal(err)

}
