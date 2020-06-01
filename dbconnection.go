package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "uv0wzcemsgoyydpa:uv0wzcemsgoyydpa@tcp(bwuogspv0wb9a6xoxprw-mysql.services.clever-cloud.com:3306)/bwuogspv0wb9a6xoxprw")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	return db
}
