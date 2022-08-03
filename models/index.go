package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var logger *log.Logger

func init() {
	currentDate := time.Now().UTC().Format("2006-Jan-2")
	filename := "../logs/DB-" + currentDate
	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	logger = log.New(f, "CMS-", log.Lshortfile)
	LogMessage("START AGAIN")

	var err error
	DB, err = sql.Open(dbConfig.dialect, dbConfig.credential)
	if err != nil {
		panic(err)
	}
}

func LogMessage(msg string) {
	logger.Println(time.Now().UTC().String() + msg)
}
