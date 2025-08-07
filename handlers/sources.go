package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// SourcesHandler handles the /sources endpoint for managing RSS feed sources
func SourcesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetSourcesHandler(w, r)
	case http.MethodPost:
		PostSourcesHandler(w, r)
	case http.MethodDelete:
		DeleteSourcesHandler(w, r)
	default:
		SourcesRes(w, http.StatusMethodNotAllowed, "Method not allowed. Use GET to retrieve sources, POST to add a source, or DELETE to remove a source.", 0, nil)
	}
}

// GetSourcesHandler handles GET requests to /sources endpoint
func GetSourcesHandler(w http.ResponseWriter, r *http.Request) {
	srcStatus := http.StatusOK
	srcMes := "Fetching RSS feed sources..."

	// Add error checking for potential server errors
	if MyBlogs == nil {
		srcStatus = http.StatusInternalServerError
		srcMes = "Internal server error: RSS sources not initialized"
		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Println("RSS sources not initialized")
		return
	}

	// Return all RSS feed sources
	srcMes = srcMes + "Fetched sources successfully"

	SourcesRes(w, srcStatus, srcMes, len(MyBlogs), MyBlogs)
}

// PostSourcesHandler handles POST requests to /sources endpoint to add a new source
func PostSourcesHandler(w http.ResponseWriter, r *http.Request) {
	srcStatus := http.StatusOK
	srcMes := "Adding RSS feed source..."

	// Check for potential server errors
	if MyBlogs == nil {
		srcStatus = http.StatusInternalServerError
		srcMes = "Internal server error: RSS sources not initialized"
		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Println("RSS sources not initialized during POST operation")
		return
	}

	// Parse the JSON request body
	var requestBody SourceRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		srcStatus = http.StatusBadRequest
		srcMes = "Invalid JSON format: " + err.Error()

		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	// Validate that source URL is not empty
	if requestBody.Source == "" {
		srcStatus = http.StatusBadRequest
		srcMes = "Source URL cannot be empty"

		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Println("Source URL is empty")
		return
	}

	// Check if source already exists
	for _, existingSource := range MyBlogs {
		if existingSource == requestBody.Source {
			srcStatus = http.StatusConflict
			srcMes = "Source already exists: " + requestBody.Source

			SourcesRes(w, srcStatus, srcMes, len(MyBlogs), MyBlogs)
			log.Printf("Source already exists: %s", requestBody.Source)
			return
		}
	}

	// Add the new source to MyBlogs
	oldCount := len(MyBlogs)
	MyBlogs = append(MyBlogs, requestBody.Source)
	newCount := len(MyBlogs)

	// Verify the operation succeeded
	if newCount != oldCount+1 {
		srcStatus = http.StatusInternalServerError
		srcMes = "Internal server error: Failed to add source to list"
		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Printf("Failed to add source to list: expected count %d, got %d", oldCount+1, newCount)
		return
	}

	srcMes = srcMes + "Successfully added source: " + requestBody.Source + ". Total sources: " + strconv.Itoa(newCount)

	SourcesRes(w, srcStatus, srcMes, newCount, MyBlogs)
	log.Printf("Source added successfully. Old count: %d, New count: %d, Added: %s", oldCount, newCount, requestBody.Source)
}

// DeleteSourcesHandler handles DELETE requests to /sources endpoint to remove a source
func DeleteSourcesHandler(w http.ResponseWriter, r *http.Request) {
	srcStatus := http.StatusOK
	srcMes := "Removing RSS feed source..."

	// Check for potential server errors
	if MyBlogs == nil {
		srcStatus = http.StatusInternalServerError
		srcMes = "Internal server error: RSS sources not initialized"
		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Println("RSS sources not initialized during DELETE operation")
		return
	}

	// Parse the JSON request body
	var requestBody SourceRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		srcStatus = http.StatusBadRequest
		srcMes = "Invalid JSON format: " + err.Error()

		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	// Validate that source URL is not empty
	if requestBody.Source == "" {
		srcStatus = http.StatusBadRequest
		srcMes = "Source URL cannot be empty"

		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Println("Source URL is empty")
		return
	}

	// Find and remove the source from MyBlogs
	oldCount := len(MyBlogs)
	sourceFound := false
	newSources := make([]string, 0, len(MyBlogs))

	for _, existingSource := range MyBlogs {
		if existingSource != requestBody.Source {
			newSources = append(newSources, existingSource)
		} else {
			sourceFound = true
		}
	}

	if !sourceFound {
		srcStatus = http.StatusNotFound
		srcMes = "Source not found: " + requestBody.Source

		SourcesRes(w, srcStatus, srcMes, len(MyBlogs), MyBlogs)
		log.Printf("Source not found: %s", requestBody.Source)
		return
	}

	// Update MyBlogs with the new list
	MyBlogs = newSources
	newCount := len(MyBlogs)

	// Verify the operation succeeded
	if newCount != oldCount-1 {
		srcStatus = http.StatusInternalServerError
		srcMes = "Internal server error: Failed to remove source from list"
		SourcesRes(w, srcStatus, srcMes, 0, nil)
		log.Printf("Failed to remove source from list: expected count %d, got %d", oldCount-1, newCount)
		return
	}

	srcMes = srcMes + "Successfully removed source: " + requestBody.Source + ". Total sources: " + strconv.Itoa(newCount)

	SourcesRes(w, srcStatus, srcMes, newCount, MyBlogs)
	log.Printf("Source removed successfully. Old count: %d, New count: %d, Removed: %s", oldCount, newCount, requestBody.Source)
}
