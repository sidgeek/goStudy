package dbops

import (
	"database/sql"

	_ "github.com/go-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123@#@tcp()/video_server?chatset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
