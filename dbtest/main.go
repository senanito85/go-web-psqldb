package main

import (
	"database/sql"
	"fmt"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1994"
	DB_NAME     = "postgres"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Inserting values")

	sum := 0
	for i := 1; i <= 50; i++ {
		sum += i

		uid := randomdata.Number(9999)
		name := randomdata.SillyName()
		mail := randomdata.Email()
		phone := randomdata.PhoneNumber()

		err = db.QueryRow("INSERT INTO userinfo(uid,name,email,phone) VALUES($1,$2,$3,$4) returning uid;", uid, name, mail, phone).Scan(&uid)
		checkErr(err)

	}

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT uid,name,email,phone FROM userinfo;")
	checkErr(err)

	for rows.Next() {
		var uid int
		var name string
		var email string
		var phone string
		err = rows.Scan(&uid, &name, &email, &phone)
		checkErr(err)
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, name, email, phone)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
