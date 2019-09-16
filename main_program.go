package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "swapnil"
	dbName := "go_db"
	dbHost := "172.18.0.2"
	dbPort := "3306"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)

	if err != ni; {
		panic(err.Error())
	}

	return db
}

var tmpl := template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")

	if err != nil {
        panic(err.Error())
    }

	emp := Employee{}
	res := []Employee}{}

	for selDB.Next(){
		var id int
		var name, city string

		err = selDB.Scan
	}
}
