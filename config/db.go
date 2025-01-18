package config

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type DataBase struct {
	dbase *sql.DB
}

func Connect() *DataBase {
	var str strings.Builder
	str.WriteString("user=" + GetFromEnv("DBUSER") + " dbname=" + GetFromEnv("DBNAME") + " password=" + GetFromEnv("PASS") + " sslmode=disable")
	db, err := sql.Open("postgres", str.String())
	if err != nil {
		log.Fatalln("Can't connect to DB PostgreSQL", err)
	}

	return &DataBase{
		dbase: db,
	}
}

func (db *DataBase) GetDB() *sql.DB {
	return db.dbase
}
