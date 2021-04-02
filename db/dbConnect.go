package db

import (
	"database/sql"
	"fmt"
	"gin-server/db"
	_ "github.com/go-sql-driver/mysql"
)

var db1 = 

//Db에 접속하여 인서트를 하는 함수.
func Connect() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/db_study")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO user (name, id, pw) VALUES ('배태현', 'baetaehyeon', '040809')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
