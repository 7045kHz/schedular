package postgres

import (
	"database/sql"
	"os"

	"github.com/7045kHz/schedular/utils"
	"github.com/lib/pq"
)

var pgdb *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("PG_CONNECT"))
	utils.LogFatal(err)

	pgdb, err = sql.Open("postgres", pgUrl)
	utils.LogFatal(err)

	err = pgdb.Ping()
	utils.LogFatal(err)

	return pgdb
}
