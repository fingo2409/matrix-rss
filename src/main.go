package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Fingo2409/matrix-rss/feed"
	"github.com/Fingo2409/matrix-rss/matrix"
)

func main() {
	feedUrlsEnv := os.Getenv("FEED_URLS")
	if feedUrlsEnv == "" {
		fmt.Println("FEED_URLS environment variable not set")
		os.Exit(1)
	}
	feedURLs := strings.Split(feedUrlsEnv, ",")

	matrixServer := os.Getenv("MATRIX_SERVER")
	matrixRoomID := os.Getenv("MATRIX_ROOM_ID")
	matrixToken := os.Getenv("MATRIX_TOKEN")
	checkIntervalStr := os.Getenv("CHECK_INTERVAL")
	if matrixServer == "" || matrixRoomID == "" || matrixToken == "" || checkIntervalStr == "" {
		fmt.Println("One or more required environment variables are not set")
		os.Exit(1)
	}
	checkInterval, err := strconv.Atoi(checkIntervalStr)
	if err != nil {
		fmt.Println("Invalid CHECK_INTERVAL value, must be an integer")
		os.Exit(1)
	}
	lastUpdates := make(map[string]string)

	for {
		for _, feedURL := range feedURLs {
			fetchedFeed, err := feed.FetchFeed(feedURL)
			if err != nil {
				fmt.Println("Error fetching feed:", err)
				continue
			}

			if len(fetchedFeed.Entries) > 0 && fetchedFeed.Entries[0].Updated != lastUpdates[feedURL] {
				lastUpdates[feedURL] = fetchedFeed.Entries[0].Updated

				message := fmt.Sprintf("IT-News: [%s](%s)", fetchedFeed.Entries[0].Title, fetchedFeed.Entries[0].Link.Href)
				err = matrix.SendMatrixMessage(matrixServer, matrixRoomID, matrixToken, message)
				if err != nil {
					fmt.Println("Error sending Matrix message:", err)
				} else {
					fmt.Println("Update message sent for feed:", feedURL)
				}
			}
		}
		time.Sleep(time.Duration(checkInterval) * time.Minute)
	}
}
