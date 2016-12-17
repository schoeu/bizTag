package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"config"
)

var db *sql.DB

func getDB() *sql.DB {
	var cfg = config.GetConf()

	tDB, err := sql.Open("mysql", cfg.DBUsername+":"+cfg.DBPassword+cfg.DBAddress+"/"+cfg.DBName)
	if err != nil {
		panic(err.Error())
	}
	db = tDB
	return db
}

func CloseDB() {
	db.Close()
}
