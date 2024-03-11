package contentscrapper

import (
	"fmt"
	newsdata "google-news-scrapper/types"
	parsehtml "google-news-scrapper/utils/parse-html"
	"net/http"

	"golang.org/x/net/html"
)

func RunContentScapper(url string) ([]newsdata.Data, error) {
	list_of_news := []newsdata.Data{}
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to make HTTP call", err)
		return list_of_news, err
	}
	defer resp.Body.Close()

	// Read the response body
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error when parsing HTTP response body:", err)
		return list_of_news, err
	}

	// make a recursive call to your function
	parsehtml.ParseBody(doc, &list_of_news)
	return list_of_news, nil
}
