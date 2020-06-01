package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/login")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	return db
}
