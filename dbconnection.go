package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "bxedpdnv1mniuhiszlm5-mysql.services.clever-cloud.com:BNNhc81iQJixAyW0aCZ6@tcp(bxedpdnv1mniuhiszlm5-mysql.services.clever-cloud.com:3306)/bxedpdnv1mniuhiszlm5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	return db
}
