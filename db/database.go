package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" 
)

const (
	host     = "localhost"
	port     = 5432 
	user     = "postgres"
	password = "root"
	dbname   = "colorpaletteexplorer"
)

var DB *sql.DB

func Init() {
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database")

}
