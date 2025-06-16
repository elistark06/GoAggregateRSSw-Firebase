package main

// Add these build tags at the top of the file

import (
	ArticleRequests "GoAggregateRSS/databaseFolder"
	serverFolder "GoAggregateRSS/serverFolder"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	_ "github.com/mmcdole/gofeed"   // Import the gofeed package for RSS parsing
)

func main() {

	s := &serverFolder.Server{
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ReadTimeout:       10000 * time.Second, // No timeout
		WriteTimeout:      10000 * time.Second, // No timeout
		IdleTimeout:       10000 * time.Second, // No timeout
		ReadHeaderTimeout: 10000 * time.Second, // No timeout
	}
	// Start the server
	go serverFolder.Start(s)

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "./articles.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Initialize the ArticleRequests struct with the database connection
	ars := &ArticleRequests.ArticleRequests{
		Db: db,
	}

	// Fetch the RSS feed and parse the articles
	articles, err := ars.FetchRSSFeed()
	if err != nil {
		fmt.Println("Error fetching RSS feed:", err)
		return
	}

	if articles == nil {
		fmt.Println("No articles found.")
		return
	} else {
		fmt.Println("Articles fetched successfully:", len(articles))
	}

	// Deprecated code for creating the table
	// Uncomment the following lines if you want to create the table in the database
	/*
		if err != nil {
			log.Fatal(err)
		}

		defer databaseConnection.Close()

		articleRepository := &database.ArticleRequests{Db: databaseConnection}


			err = articleRepository.CreateTable()
			if err != nil {
				log.Fatal(err, "ERROR WITH TABLE")
			}
	*/

}
