//
//
// Commented out to avoid linting issues or accidental execution, this is to WIPE THE DATABASE only.
//
//
/*

package main

import (
	ArticleRequests "GoAggregateRSS/databaseFolder"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database connection
	db, err := sql.Open("sqlite3", "./articles.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create ArticleRequests instance
	ars := &ArticleRequests.ArticleRequests{
		Db: db,
	}

	// Use ArticleRequests to wipe database
	err = wipeDatabase(ars)
	if err != nil {
		log.Fatal("Error wiping database:", err)
	}
}

func wipeDatabase(ars *ArticleRequests.ArticleRequests) error {
	// Drop the existing table
	_, err := ars.Db.Exec(`DROP TABLE IF EXISTS articles`)
	if err != nil {
		return fmt.Errorf("error dropping table: %v", err)
	}

	// Create a new table
	_, err = ars.Db.Exec(`CREATE TABLE articles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        link TEXT,
        content TEXT,
        publisheddate TEXT,
        receiveddate TEXT
    )`)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	fmt.Println("Database wiped successfully!")
	return nil
}
*/