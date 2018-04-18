package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
        "log"
)

func test_db() {
	db, err := sql.Open("mysql",
		"django:djangopass@tcp(127.0.0.1:3306)/homepage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // danger!
	var (
         table string
        )
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(table)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
