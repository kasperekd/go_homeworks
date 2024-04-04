package main

import (
	"database/sql"
	dbPack "example_mock/internal/db"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=go_test password=go_test dbname=go_test_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		log.Println(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
