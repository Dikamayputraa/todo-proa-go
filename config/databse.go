package config

import (
	"database/sql"
	_ "mysql-master"
)

func DBconn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todo")
	
	return db, err
}
