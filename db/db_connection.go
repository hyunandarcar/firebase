package db

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Go ..........1")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/rest_database")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}
