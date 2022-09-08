package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Implementation of data:
	insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Alex', 25)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Println("MySQL is connected")
}
