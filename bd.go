package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "03795"
	dbname   = "users"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// insert, err := db.Query(`INSERT INTO "users" ("name", "age") VALUES('Alexxxx', '12')`)
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	res, err := db.Query("SELECT name, age FROM users")

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
		//fmt.Println(1)
	}
	fmt.Println("Successfully connected!")
}
