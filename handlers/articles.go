package handlers

import (
	"context"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
)

// ArticlesHandler handles the /articles endpoint
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var articleArray map[string]Article

	// Create a reference to the articles node in the database
	articlesRef := FbDB.NewRef("articles")
	// Fetch all articles
	err := articlesRef.Get(context.Background(), &articleArray)

	arsStatus := http.StatusOK
	arsMes := "Fetching articles..."

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		arsStatus = http.StatusMethodNotAllowed
		arsMes = "Method not allowed"

		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// Check if articles ref throws an error
	if err != nil {
		arsStatus = http.StatusInternalServerError
		arsMes = "Error fetching articles: " + err.Error()
		log.Println("Error fetching articles:", err)

		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// if the articles array has items then we can return them
	if len(articleArray) > 0 {
		arsMes = arsMes + "Fetched articles successfully"

		ArticlesRes(w, arsStatus, arsMes, len(articleArray), articleArray)
		return
	}

	if len(articleArray) == 0 {
		// if the articles array is empty, return a 404
		arsStatus = http.StatusNotFound
		log.Println("No articles found")

		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}
}

// LatestArticlesHandler handles the /articles/latest endpoint
func LatestArticlesHandler(w http.ResponseWriter, r *http.Request) {
	// default amount
	amount := 10

	// Check if a query parameter is provided
	amountParam := r.URL.Query().Get("amount")

	if amountParam != "" {
		if num, err := strconv.Atoi(amountParam); err == nil {
			amount = num
			log.Println("Amount set to:", amount)
		}
	} else {
		log.Println("No amount parameter provided, defaulting to 10")
	}

	var articleArray map[string]Article

	// Create a reference to the articles node in the database
	articlesRef := FbDB.NewRef("articles")
	// Fetch the latest articles ordered by publishedDate
	err := articlesRef.OrderByChild("publishedDate").LimitToLast(amount).Get(context.Background(), &articleArray)

	arsStatus := http.StatusOK
	arsMes := "Fetching articles..."

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		arsStatus = http.StatusMethodNotAllowed
		arsMes = "Method not allowed"

		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// Check if articles ref throws an error
	if err != nil {
		arsStatus = http.StatusInternalServerError
		arsMes = "Error fetching articles: " + err.Error()
		log.Println("Error fetching articles:", err)
		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// if the articles array has items then we can return them
	if len(articleArray) > 0 {
		arsMes = arsMes + "Fetched articles successfully"
		ArticlesRes(w, arsStatus, arsMes, len(articleArray), articleArray)
		return
	}
}

// RandomArticlesHandler handles the /articles/rand endpoint
func RandomArticlesHandler(w http.ResponseWriter, r *http.Request) {
	arsMes := "Fetching random articles..."
	arsStatus := http.StatusOK

	// Sets function variables to hold response while the proccess runs
	var articleArray map[string]Article

	randomArticles := make(map[string]Article)

	// Create a reference to the articles node in the database
	articlesRef := FbDB.NewRef("articles")

	amount := 10 // default amount

	// Check if a query parameter is provided
	amountParam := r.URL.Query().Get("amount")
	if amountParam != "" {
		if num, err := strconv.Atoi(amountParam); err == nil {
			amount = num
			log.Println("Amount set to:", amount)
		} else {
			log.Println("Invalid amount parameter, defaulting to 10")
		}

	} else {
		log.Println("No amount parameter provided, defaulting to 10")
		arsMes = arsMes + " No amount parameter provided, defaulting to 10. To get a specific amount, use ?amount=1 (<- use your value instead of 1) in the URL."
	}

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		arsStatus = http.StatusMethodNotAllowed
		arsMes = "Method not allowed"

		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// Fetch all articles
	err := articlesRef.Get(context.Background(), &articleArray)

	// Check if articles ref throws an error
	if err != nil {
		arsStatus = http.StatusInternalServerError
		arsMes = "Error fetching articles: " + err.Error()
		log.Println("Error fetching articles:", err)

		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// if the articles array has items then we can return them
	if len(articleArray) > 0 {
		arsMes = arsMes + "Fetched articles successfully"

		// Get the keys of the articles
		keys := make([]string, 0, len(articleArray))
		for key := range articleArray {
			keys = append(keys, key)
		}

		// Shuffle the keys to get random articles
		rand.Shuffle(len(keys), func(i, j int) {
			keys[i], keys[j] = keys[j], keys[i]
		})

		// Select the first 'amount' keys
		for i := 0; i < amount && i < len(keys); i++ {
			randomIndex := i
			randomArticles[keys[randomIndex]] = articleArray[keys[randomIndex]]
		}

		ArticlesRes(w, arsStatus, arsMes, len(randomArticles), randomArticles)
		return
	}
}

// SearchArticlesHandler handles the /articles/search endpoint for keyword search
func SearchArticlesHandler(w http.ResponseWriter, r *http.Request) {
	// Sets function variables to hold response while the proccess runs
	var articleArray map[string]Article
	filteredArticles := make(map[string]Article)

	// Create a reference to the articles node in the database
	articlesRef := FbDB.NewRef("articles")
	// Fetch all articles
	err := articlesRef.Get(context.Background(), &articleArray)

	keyword := "cloud" // default keyword

	// Check if a query parameter is provided
	keywordParam := r.URL.Query().Get("keyword")

	arsStatus := http.StatusOK
	arsMes := "Searching for articles..."

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		arsStatus = http.StatusMethodNotAllowed
		arsMes = "Method not allowed"
		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	if keywordParam != "" {
		keyword = keywordParam
		log.Println("Searching for keyword:", keyword)
	} else {
		log.Println("No keyword parameter provided, defaulting to 'cloud'")
		arsMes = arsMes + " No keyword parameter provided, defaulting to 'cloud'. To search for a specific keyword, use ?keyword=value in the URL."
	}

	// Check if articles ref throws an error
	if err != nil {
		arsStatus = http.StatusInternalServerError
		arsMes = "Error fetching articles: " + err.Error()
		log.Println("Error fetching articles:", err)
		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}

	// Filter articles by keyword in title (case insensitive)
	if len(articleArray) > 0 {
		// Loop through each article in the articleArray search for the keyword
		for key, article := range articleArray {

			// Convert both title and keyword to lowercase for case-insensitive search
			titleLower := strings.ToLower(article.Title)
			keywordLower := strings.ToLower(keyword)

			// Check if the title contains the keyword
			if strings.Contains(titleLower, keywordLower) {
				filteredArticles[key] = article
			}
		}

		// if filtered articles found
		if len(filteredArticles) > 0 {
			arsMes = arsMes + " Found " + strconv.Itoa(len(filteredArticles)) + " articles matching keyword: " + keyword
			ArticlesRes(w, arsStatus, arsMes, len(filteredArticles), filteredArticles)
			return
		}
	}

	if len(filteredArticles) == 0 {
		// if no matching articles found, return a 404
		arsStatus = http.StatusNotFound
		arsMes = arsMes + " No articles found matching keyword: " + keyword
		ArticlesRes(w, arsStatus, arsMes, 0, nil)
		return
	}
}
