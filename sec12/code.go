package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	var cmd string
	var err error

	// cmd = `create table if not exists person(
	//        name string,
	//        age int
	//    )`
	// _, err := DbConnection.Exec(cmd)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// insert
	// cmd = "insert into person (name, age) values (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Nancy", "20")
	// _, err = DbConnection.Exec(cmd, "Mike", "24")

	// update
	// cmd = "update person set age = ? where name = ?"
	// _, err = DbConnection.Exec(cmd, "22", "Nancy")

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// multiple select
	cmd = "select * from person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var people []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)

		if err != nil {
			log.Fatalln(err)
		}
		people = append(people, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range people {
		fmt.Println(p)
	}

	// single select
	cmd = "select * from person where name = ?"
	row, _ := DbConnection.Query(cmd, "Nancy")
	var p Person
	row.Next()
	err = row.Scan(&p.Name, &p.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {

			log.Fatalln(err)
		}
	}
	fmt.Println(p)

	fmt.Println("success.")
}
