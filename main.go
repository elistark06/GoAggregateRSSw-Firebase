package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	databaseConnection, err := sql.Open("sqlite3", "./articles.db")

	if err != nil {
		log.Fatal(err)
	}

	defer databaseConnection.Close()

	articleRepository := &database.ArticleRepository{Db: databaseConnection}

	err = articleRepository.CreateTable()
	if err != nil {
		log.Fatal(err, "ERROR WITH TABLE")
	}
}
