package schedule

import (
	"database/sql"
	"fmt"

	"github.com/7045kHz/schedular/models"
	"github.com/7045kHz/schedular/utils"
)

func GetAllJobs(mssqldb *sql.DB) (jobs []models.Jobs) {
	sqlSelect := `SELECT [Id]
	,[Name]
	,[Execution_Server]
	,[Enabled]
	,[Job_Definition]
	,[Job_Schedule]
	,[Created_On]
	,[Last_Updated]
	,[Started]
	,[Finished]
FROM [OSDISCOVERY].[dbo].[JOB_INVENTORY]`
	stmt, err := mssqldb.Prepare(sqlSelect)
	utils.LogFatal(err)

	defer stmt.Close()
	rows, err := stmt.Query()
	if err == sql.ErrNoRows {
		fmt.Println("No rows found")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var j models.Jobs
		_ = rows.Scan(&j.Job_Id, &j.Name, &j.Execution_Server, &j.Enabled, &j.Job_Definition, &j.Job_Schedule, &j.Created_On, &j.Last_Updated, &j.Started, &j.Finished)
		jobs = append(jobs, j)
	}
	return jobs
}
func GetEnabledJobs(mssqldb *sql.DB) (jobs []models.Jobs) {
	sqlSelect := `SELECT [Id]
	,[Name]
	,[Execution_Server]
	,[Enabled]
	,[Job_Definition]
	,[Job_Schedule]
	,[Created_On]
	,[Last_Updated]
	,[Started]
	,[Finished]
FROM [OSDISCOVERY].[dbo].[JOB_INVENTORY] WHERE ENABLED=1`
	stmt, err := mssqldb.Prepare(sqlSelect)
	utils.LogFatal(err)

	defer stmt.Close()
	rows, err := stmt.Query()
	if err == sql.ErrNoRows {
		fmt.Println("No rows found")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var j models.Jobs
		_ = rows.Scan(&j.Job_Id, &j.Name, &j.Execution_Server, &j.Enabled, &j.Job_Definition, &j.Job_Schedule, &j.Created_On, &j.Last_Updated, &j.Started, &j.Finished)
		jobs = append(jobs, j)
	}
	return jobs
}
func GetNowJobs(mssqldb *sql.DB, now string) (jobs []models.Jobs) {

	sqlSelect := fmt.Sprintf("SELECT [Id] ,[Name],[Execution_Server],[Enabled],[Job_Definition],[Job_Schedule],	[Created_On],[Last_Updated],[Started],[Finished] FROM [OSDISCOVERY].[dbo].[JOB_INVENTORY] WHERE ENABLED=1 and Job_Schedule like '%c%s%c' and Started not like '%c%s%c' ", 37, now, 37, 37, now, 37)
	stmt, err := mssqldb.Prepare(sqlSelect)
	utils.LogFatal(err)
	//fmt.Printf("sqlSelect = %s\n", sqlSelect)
	defer stmt.Close()
	rows, err := stmt.Query()
	if err == sql.ErrNoRows {
		fmt.Println("No rows found")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var j models.Jobs
		_ = rows.Scan(&j.Job_Id, &j.Name, &j.Execution_Server, &j.Enabled, &j.Job_Definition, &j.Job_Schedule, &j.Created_On, &j.Last_Updated, &j.Started, &j.Finished)
		jobs = append(jobs, j)
	}
	return jobs
}
