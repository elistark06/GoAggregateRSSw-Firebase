package main

import (
	"log"
	"net/http"

	"RSSAggregator/handlers"
)

// rootHandler handles the root endpoint
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Welcome to the RSS Aggregator!"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {

	MyBlogs := []string{
		"https://us2.campaign-archive.com/feed?u=5005148108dfbac726f74e31e&id=239e48d26e",
		"https://www.youtube.com/feeds/videos.xml?channel_id=UCH7xyou6RXO8PKwMZ4nQ64Q",
		"https://cloudblog.withgoogle.com/products/networking/rss/",
		"https://feeds.feedburner.com/WebDeveloperJuice",
		"https://techcrunch.com/tag/saas/feed/",
		"https://news.google.com/rss/search?hl=en-US&gl=US&q=startups&um=1&ie=UTF-8&ceid=US:en",
		"https://www.cloudcomputing-news.net/feed/",
		"https://feeds.feedburner.com/cioreview/fvHK",
		"https://www.cloudally.com/feed/",
		"https://www.wired.com/feed/tag/ai/latest/rss",
		"https://www.techrepublic.com/rssfeeds/topic/cloud/",
		"https://www.techrepublic.com/rssfeeds/topic/cloud-security/",
		"https://www.techrepublic.com/rssfeeds/topic/google/",
	}

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/articles", handlers.ArticlesHandler)

	http.HandleFunc("/sources", handlers.SourcesHandler)

	log.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
