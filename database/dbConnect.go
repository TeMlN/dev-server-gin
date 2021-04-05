package database

import (
	"database/sql"
	"fmt"
	"gin-server/config"

	_ "github.com/go-sql-driver/mysql"
)

//Db에 접속하여 인서트를 하는 함수.
func Connect() {
	dbConfig := config.Conf.Database

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Server,
		dbConfig.Port,
		dbConfig.Database)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var count int

	err = db.QueryRow("SELECT 1 + 1").Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

}
