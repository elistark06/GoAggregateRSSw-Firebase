package database

import (
	"database/sql"
)

type ArticleRepository struct {
	Db *sql.DB
}

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
