package server

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB{
    var cfg = getConf()

    db, err := sql.Open("mysql", cfg.DBUsername + ":" + cfg.DBPassword + cfg.DBAddress + "/" + cfg.DBName)
    if err != nil {
        panic(err.Error())  
    }
    return db
}