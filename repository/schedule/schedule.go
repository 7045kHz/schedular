package schedule

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/7045kHz/schedular/models"
	"github.com/7045kHz/schedular/utils"
)

func GetAllJobs(mssqldb *sql.DB) (jobs []models.Jobs) {
	var inventory_table = os.Getenv("INVENTORY_TABLE")
	sqlSelect := fmt.Sprintf(`SELECT [Id],[Name],[Execution_Server],[Enabled],[Job_Definition],[Job_Schedule],[Created_On],[Last_Updated],[Started],[Finished] FROM %s`, inventory_table)
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

// func GetMyJobs(mssqldb *sql.DB) (jobs []models.Jobs) {
// 	hostname, err := utils.GetHostname()
// 	utils.LogFatal(err)
// 	var inventory_table = os.Getenv("INVENTORY_TABLE")
// 	sqlSelect := fmt.Sprintf(`SELECT [Id],[Name],[Execution_Server],[Enabled],[Job_Definition],[Job_Schedule],[Created_On],[Last_Updated],[Started],[Finished] FROM %s WHERE Execution_Server=%s`, inventory_table, hostname)
// 	stmt, err := mssqldb.Prepare(sqlSelect)
// 	utils.LogFatal(err)

// 	defer stmt.Close()
// 	rows, err := stmt.Query()
// 	if err == sql.ErrNoRows {
// 		fmt.Println("No rows found")
// 		return nil
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var j models.Jobs
// 		_ = rows.Scan(&j.Job_Id, &j.Name, &j.Execution_Server, &j.Enabled, &j.Job_Definition, &j.Job_Schedule, &j.Created_On, &j.Last_Updated, &j.Started, &j.Finished)
// 		jobs = append(jobs, j)
// 	}
// 	return jobs
// }
func GetNowJobs(mssqldb *sql.DB, now string) (jobs []models.Jobs) {
	hostname, err := utils.GetHostname()
	utils.LogFatal(err)
	fmt.Printf("Hostname: %s\n", hostname)
	var inventory_table = os.Getenv("INVENTORY_TABLE")
	//sqlTest := fmt.Sprintf("UPDATE %s SET [LAST_UPDATED]=CURRENT_TIMESTAMP  ", inventory_table)

	//fmt.Println(sqlTest)
	sqlSelect := fmt.Sprintf("SELECT [Id] ,[Name],[Execution_Server],[Enabled],[Job_Definition],[Job_Schedule],	[Created_On],[Last_Updated],[Started],[Finished] FROM  %s WHERE ENABLED=1 and Job_Schedule like '%c%s%c' and (Started not like '%c%s%c' or Started is null ) and Execution_Server='%s'", inventory_table, 37, now, 37, 37, now, 37, hostname)
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
