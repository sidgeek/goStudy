package dbops

import (
	"database/sql"

	_ "github.com/go-driver/mysql"
)

func openConn() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:123@#@tcp()/video_server?chatset=utf8")
	if err != nil {
		panic(err.Error())
	}

	return dbConn
}

func AddUserCredential(loginName string, pwd string) error {
	db := openConn()
}

func GetUserCredential(loginName string) (string error) {
	db := openConn()
}
