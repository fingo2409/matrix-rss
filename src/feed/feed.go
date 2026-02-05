package feed

import (
	"encoding/xml"
	"net/http"
)

type Entry struct {
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Link    struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

func FetchFeed(url string) (*Feed, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var feed Feed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, err
	}

	return &feed, nil
}
