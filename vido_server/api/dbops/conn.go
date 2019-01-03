package dbops

import "database/sql"

var(
	dbConn *sql.DB
	err error
)

func init (){
	dbConn, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

}
