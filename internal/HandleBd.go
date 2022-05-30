package handlerBd

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func Get() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PASSWORD"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	result, err := db.Query("select * from habr_post")
	if err != nil {
		panic(err)
	}
	for result.Next() {
		var name string
		fmt.Println(name)
	}
	defer db.Close()
}
