package mssql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/7045kHz/schedular/utils"
	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB() *sql.DB {
	sqlUri := os.Getenv("SQL_CONNECT")

	mssqldb, err := sql.Open("mssql", sqlUri)
	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
	}

	err = mssqldb.Ping()
	utils.LogFatal(err)

	return mssqldb
}
