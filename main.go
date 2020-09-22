package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "admin"
	dbname = "test"
)




func main() {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psql)

	db, err := sql.Open("postgres", psql)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	insertStmt := `insert into "grpc-app"("Name", "age") values('John', 1)`
	_, e := db.Exec(insertStmt)
	if e != nil {
		log.Fatal(e)
	}

	getdata, err := db.Query(`SELECT "age", "Name"  FROM "grpc-app"`)
	defer getdata.Close()

	if err != nil {
		log.Fatal(err)
	}
	for getdata.Next() {
		var name string
		var age int

		err = getdata.Scan(&age, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(age, name)
	}


}
