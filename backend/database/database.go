package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDatabase() error {
	DB_PATH := "./test.db"

	db, err := sql.Open("sqlite3", DB_PATH)

	if err != nil {
		fmt.Println("Failed to open sqlite db ", err)
		return err
	}

	fmt.Println("Connected to db")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS 'visitors' (
		"uuid" TEXT PRIMARY KEY,
		"previous" INTEGER NOT NULL,
		"current" INTEGER NOT NULL
		)
	`)

	if err != nil {
		fmt.Println("create table failed ", err)
		return err
	}

	DB = db

	return nil
}
