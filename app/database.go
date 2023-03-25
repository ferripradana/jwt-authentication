package app

import (
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {
	host := os.Getenv("DATABASE_HOST")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+name+"?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
