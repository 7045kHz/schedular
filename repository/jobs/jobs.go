package jobs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/7045kHz/schedular/models"
)

func GetAllJobs(mssqldb *sql.DB) (jobs []models.Jobs) {
	sqlSelect := `SELECT [Id]
	,[Name]
	,[Execution_Server]
	,[Enabled]
	,[Job_Definition]
	,[Days]
	,[Hour]
	,[Minute]
	,[Created_On]
	,[Last_Updated]
	,[Last_Run]
FROM [OSDISCOVERY].[dbo].[JOB_INVENTORY]`
	stmt, err := mssqldb.Prepare(sqlSelect)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err == sql.ErrNoRows {
		fmt.Println("No rows found")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var j models.Jobs
		_ = rows.Scan(&j.Job_Id, &j.Name, &j.Execution_Server, &j.Enabled, &j.Job_Definition, &j.Days, &j.Hour, &j.Minute, &j.Created_On, &j.Last_Updated, &j.Last_Run)
		jobs = append(jobs, j)
	}
	return jobs
}
