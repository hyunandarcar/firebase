package db

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Go ..........1")

	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/rest_database")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

}
