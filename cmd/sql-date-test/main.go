package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/peterstace/date"
)

func main() {
	const connStr = "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS foo ( d DATE NOT NULL)`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO foo (d) VALUES ($1)`, date.FromTime(time.Now())); err != nil {
		log.Fatal(err)
	}

	var result date.Date
	if err := db.QueryRow("SELECT * FROM foo LIMIT 1").Scan(&result); err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	if _, err := db.Exec(`DROP TABLE IF EXISTS foo`); err != nil {
		log.Fatal(err)
	}
}
