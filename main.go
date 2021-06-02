// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1994"
	DB_NAME     = "postgres"
)

func handler(w http.ResponseWriter, r *http.Request) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	rows, err := db.Query("SELECT uid,name,email,phone FROM userinfo;")
	checkErr(err)

	for rows.Next() {
		var uid int
		var name string
		var email string
		var phone string
		err = rows.Scan(&uid, &name, &email, &phone)
		checkErr(err)
		fmt.Fprintf(w, "%3v | %8v | %6v | %6v\n", uid, name, email, phone)
	}

}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
