package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "userr14"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Berhasil Koneksi ke database")

	createTableSQL := `CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	insertSampleData(db)
}

func insertSampleData(db *sql.DB) {
	sampleUsers := []User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}

	for _, user := range sampleUsers {
		insertUserSQL := `INSERT INTO users (name, email) VALUES ($1, $2)`
		statement, err := db.Prepare(insertUserSQL)
		if err != nil {
			log.Fatal(err)
		}
		defer statement.Close()

		_, err = statement.Exec(user.Name, user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Insert data berhasil")

}
