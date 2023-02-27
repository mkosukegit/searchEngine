package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DbConfig interface {
	DbConnect()
}

var Db *sql.DB

func DbConnect() {

	connect := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))

	var err error
	Db, err = sql.Open(os.Getenv("DBMS"), connect)
	if err != nil {
		log.Println(err)
		return
	}

	log.Print("connect DB ...")
}
