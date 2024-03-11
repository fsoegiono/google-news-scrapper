package urlscrapper

import (
	"fmt"
	parseurl "google-news-scrapper/utils/parse-url"
	"net/http"

	"golang.org/x/net/html"
)

func RunURLScapper(url string) ([]string, error) {
	// Make the GET request
	list_of_url := []string{}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to make HTTP call", err)
		return list_of_url, err
	}
	defer resp.Body.Close()

	// Read the response body
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error when parsing HTTP response body:", err)
		return list_of_url, err
	}

	// make a recursive call to your function
	parseurl.ParseBody(doc, &list_of_url)
	return list_of_url, nil
}
