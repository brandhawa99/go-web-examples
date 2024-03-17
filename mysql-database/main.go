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
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
func main() {

	godotenv.Load()
	userName := goDotEnvVariable("PGUSER")
	pass := goDotEnvVariable("PGPASSWORD")
	connStr := fmt.Sprintf("postgresql://%s:%s@ep-damp-frost-a61hgvof.us-west-2.aws.neon.tech/test-mysql?sslmode=require", userName, pass)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var version string
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("version=%s\n", version)
	if err != nil {
		fmt.Println(err.Error())
	}
	query := `
		CREATE TABLE users(
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	username := "johndoe"
	password := "secret"
	createAt := time.Now()

	result, err := db.Exec(`INSERT INTO users(username, password, create_at) VALUES (?,?,?)`, username, password, createAt)
	if err != nil {
		fmt.Println(err.Error())
	}

	userId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(userId)
}
