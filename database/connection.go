package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "C:\\Users\\Adm\\Documents\\projetos\\db\\database.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
