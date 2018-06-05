package entities

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var mydb *sql.DB

func init() {
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/youhome?charset=utf8&parseTime=true")
	db, err := sql.Open("mysql", "root:root@tcp(123.207.55.27:3306)/youhome?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
