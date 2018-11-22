package main

import (
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
	"log"
)

func main() {

	// Connect string format: [username]/[password]@//[hostname]:[port]/[DB name]
	db, err := sql.Open("goracle", "bgweb/bgweb#1@//tv-uat62-dq.tvsit.co.th:1521/UAT62")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT 'Hello World!' FROM dual")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var strVal string
	for rows.Next() {
		err := rows.Scan(&strVal)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strVal)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
