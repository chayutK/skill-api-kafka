package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Sync() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		fmt.Printf("Error while connnecting to database : %s", err.Error())
		log.Fatal()
	}
	fmt.Println("------------- Database connect successfully ------------")
	return db
}
