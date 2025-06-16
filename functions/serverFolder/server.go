package serverFolder

import (
	ArticleRequests "GoAggregateRSS/databaseFolder"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mmcdole/gofeed" // Import the gofeed package for RSS parsing

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	// Import the gofeed package for RSS parsing
)

// Server struct to server metadata
type Server struct {
	Addr              string
	Handler           http.Handler
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
}

func Start(s *Server) {

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "./articles.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Set up handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		basicReply(w, http.StatusOK)
	})
	mux.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		returnAllArticles(w, r)
	})

	server := &http.Server{
		Addr:              s.Addr,
		Handler:           http.DefaultServeMux,
		ReadTimeout:       s.ReadTimeout,
		WriteTimeout:      s.WriteTimeout,
		IdleTimeout:       s.IdleTimeout,
		ReadHeaderTimeout: s.ReadHeaderTimeout,
	}

	// Create channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

}

func basicReply(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Write([]byte("help")) // Add a return value

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	//Fetch from db
	art := &ArticleRequests.ArticleRequests{}

	articles, err := art.GetArticles()
	if err != nil {
		http.Error(w, "Error fetching articles: "+err.Error(), http.StatusInternalServerError)
		log.Println("Error fetching articles:", err)
		return
	}
	if articles == nil {
		w.Write([]byte("No articles found."))
		return
	}
	for _, article := range articles {
		w.Write([]byte(article.Title + "\n")) // Write each article title to the response
	}
	log.Println("Articles returned successfully")
}
