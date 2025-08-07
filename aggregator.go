package main

import (
	"encoding/json"
	"io"
	"log"
	"time"
)

func fetchFeeds() {

	r := "./aggregateTimer.json"

	currentTime := time.Now()

	decoder := json.NewDecoder(io.Reader)

}

// parseTime converts date string to unix milliseconds
func parseTime(dateStr string) int64 {
	if dateStr == "" {
		return time.Now().UnixMilli() // Return current time if dateStr is empty
	}

	// Try multiple formats since RSS feeds can use different date formats
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		time.RFC822,
		time.RFC822Z,
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02 15:04:05 -0700 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"02 Jan 2006 15:04 MST",
		"02 Jan 2006 15:04:05 MST",
	}

	for _, format := range formats {
		// Try to parse the date string with the current format
		t, err := time.Parse(format, dateStr)

		// If parsing is successful, return the time in milliseconds since epoch
		if err == nil {
			return t.UnixMilli() // Return the time in milliseconds since epoch
		}
	}

	log.Println("Failed to parse date:", dateStr, "with all formats. Returning current time.")
	return time.Now().UnixMilli() // Return current time if all parsing attempts fail
}
