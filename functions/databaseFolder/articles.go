package database

import (
	"database/sql"
	"fmt"

	"github.com/mmcdole/gofeed"
)

// simple struct for making requests and updating the database
type ArticleRequests struct {
	Db *sql.DB
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

func (r *ArticleRequests) CreateTable() error {
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

		fp := gofeed.NewParser()     // creates a parser
		feed, _ := fp.ParseURL(blog) // fetches the feed from the URL and parses it
		if feed == nil {
			fmt.Println("No feed found for URL:", blog)
			continue
		}

		if len(feed.Items) > 0 { // checks if the feed has items
			for _, item := range feed.Items {

				// create an article for each item in the feed
				//  and pass the data to a function that will
				//  add it to the database

				Article := Article{
					Title:         item.Title,
					Link:          item.Link,
					Content:       item.Content,
					PublishedDate: item.Published,
					ReceivedDate:  item.Updated,
				}

				fmt.Println("Article found:", Article.Title)

				// Insert the article into the database
				err := r.InsertArticle(Article)
				if err != nil {
					fmt.Println("Error inserting article:", err)
					continue
				} else {
					fmt.Println("Article inserted successfully:", Article.Title)
				}

			}
		}

	}

	return nil, nil
}

func (r *ArticleRequests) InsertArticle(article Article) error {
	var count int
	// Check if the article already exists in the database
	err := r.Db.QueryRow(`
		SELECT COUNT(*) FROM articles
        WHERE title = ? AND link = ?`,
		article.Title, article.Link).Scan(&count)

	if err != nil {
		return fmt.Errorf("error checking if article exists: %v", err)
	}

	if count > 0 {
		fmt.Printf("Skipping duplicate article: %s\n", article.Title)
		return nil
	}

	// Insert the article into the database
	result, err := r.Db.Exec(`
        INSERT INTO articles (title, link, content, publisheddate, receiveddate)
        VALUES (?, ?, ?, ?, ?)
    `, article.Title, article.Link, article.Content, article.PublishedDate, article.ReceivedDate)
	if err != nil {
		return fmt.Errorf("insert failed: %v", err)
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %v", err)
	}

	// Update the article's ID
	article.ID = id
	fmt.Printf("Article inserted with ID: %d - Title: %s\n", id, article.Title)

	return nil
}
