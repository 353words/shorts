package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := "8fa74e2c92ad46f2923e9d9c1ed468e3"
	query := `
	SELECT id, name, email
	FROM users
	WHERE id=@id
	`
	row := db.QueryRow(query, sql.Named("id", id))
	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
