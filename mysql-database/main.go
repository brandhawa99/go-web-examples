package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	pgusername := goDotEnvVariable("PGUSER")
	pgpassword := goDotEnvVariable("PGPASSWORD")

	connStr := fmt.Sprintf("postgresql://%s:%s@ep-long-hill-a64guqb7.us-west-2.aws.neon.tech/test?sslmode=require", pgusername, pgpassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	// query := `
	// 	CREATE TABLE users (
	// 	  id SERIAL,
	//  	 	username TEXT NOT NULL,
	//   	password TEXT NOT NULL,
	//   	created_at TIMESTAMP,
	//   	PRIMARY KEY (id)
	// );`
	// _, err = db.Exec(query)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	{
		//inserting a new row
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)`, username, password, createdAt)
		if err != nil {
			fmt.Println(err.Error())
		}

		if err != nil {
			fmt.Println(err.Error())
		}

	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = $1"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
