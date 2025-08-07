package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

// Article represents an RSS article
type Article struct {
	Title        string `json:"title"`
	Link         string `json:"link"`
	Content      string `json:"content"`
	ReceivedDate int64  `json:"receivedDate"`
	Source       string `json:"source"`
}

// SourceRequest represents a request to add/remove RSS feed sources
type SourceRequest struct {
	Source string `json:"source"`
}

// Shared global variables (exported to be accessible from other packages)
var (
	FbApp   *firebase.App
	FbDB    *db.Client
	MyBlogs []string
)

// ArticlesRes function, takes arguments from handlers to send a JSON response for articles
func ArticlesRes(w http.ResponseWriter, status int, message string, count int, articles map[string]Article) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   status,
		"message":  message,
		"count":    count,
		"articles": articles,
	})
	log.Printf("Response sent with status: %d, message: %s, count: %d", status, message, count)
}

// SourcesRes function, takes arguments from handlers to send a JSON response for sources
func SourcesRes(w http.ResponseWriter, status int, message string, count int, sources []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"message": message,
		"count":   count,
		"sources": sources,
	})
	log.Printf("Response sent with status: %d, message: %s, count: %d", status, message, count)
}
