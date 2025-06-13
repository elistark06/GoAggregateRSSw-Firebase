package database

import (
	"database/sql"
	"fmt"

	"github.com/mmcdole/gofeed"
	_ "github.com/mmcdole/gofeed"
)

// responsible for database interactions
type ArticleRepository struct {
	Db *sql.DB
}

// simple struct for request methods
type ArticleRequests struct {
}

// assigns fields for article
type Article struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Link          string `json:"link"`
	Content       string `json:"content"`
	PublishedDate string `json:"published_date"`
	ReceivedDate  string `json:"received_date"`
}

func (r *ArticleRepository) CreateTable() error {
	_, err := r.Db.Exec(`CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title STRING,
		link STRING,
		content STRING,
		publisheddate STRING,
		receiveddate STRING
	)`)

	return err
}

func (r *ArticleRequests) FetchRSSFeed() ([]Article, error) {
	myBlogs := []string{
		"https://us2.campaign-archive.com/feed?u=5005148108dfbac726f74e31e&id=239e48d26e",
		"https://www.youtube.com/feeds/videos.xml?channel_id=UCH7xyou6RXO8PKwMZ4nQ64Q",
		"https://cloudblog.withgoogle.com/products/networking/rss/",
		"https://feeds.feedburner.com/WebDeveloperJuice",
		"https://techcrunch.com/tag/saas/feed/",
		"https://news.google.com/rss/search?hl=en-US&gl=US&q=startups&um=1&ie=UTF-8&ceid=US:en",
		"https://www.cloudcomputing-news.net/feed/",
		"https://feeds.feedburner.com/cioreview/fvHK",
		"https://www.cloudally.com/feed/",
	}

	for _, blog := range myBlogs {

		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(blog)
		if feed == nil {
			fmt.Println("No feed found for URL:", blog)
			continue
		}

		if len(feed.Items) > 0 {
			fmt.Println("first result", feed.Items)
		}
	}

	return nil, nil
}

func (r *ArticleRepository) InsertArticle(article Article) error {
