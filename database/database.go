package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/ads_indonesia?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(10000)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Minute)
	return db
}
