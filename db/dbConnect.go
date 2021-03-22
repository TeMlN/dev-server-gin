package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"]]]''
	]]\\]\]
	ioiiuuuuuuhfmt vg5o
)

func Connect() {
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
     
    //...(db 사용)....
}