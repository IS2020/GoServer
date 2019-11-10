package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
var apis = make(map[string]int)
var cur *sql.DB
func startSql(){
	db, err := sql.Open("mysql", "root@/GesEventos")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	cur = db
}
func loadApis()  {
	var col1 int
	var col2 []byte
	rows, err := cur.Query("SELECT id_antena,api_key FROM Antenas")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for rows.Next() {
		// Scan the value to []byte
		err = rows.Scan(&col1, &col2)

		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}

		// Use the string value
		apis[string(col2)] = col1
	}
	fmt.Println(apis)
}
func newReport(id_antena int, filepath string,t time.Time){
	stmtIns, err := cur.Prepare("INSERT INTO Reportes VALUES( ?, ? ,?,?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(nil,id_antena,filepath,t) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}