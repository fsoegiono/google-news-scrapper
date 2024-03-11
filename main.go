package main

import (
	"fmt"
	contentscrapper "google-news-scrapper/libs/features/content-scrapper"
	urlscrapper "google-news-scrapper/libs/features/url-scrapper"
	newsdata "google-news-scrapper/types"
)

func main() {
	data := []newsdata.Data{}
	urls, error := urlscrapper.RunURLScapper("https://news.google.com/home?hl=id&gl=ID&ceid=ID%3Aid")

	if error != nil {
		return
	}

	for _, url := range urls {
		newsdata, error := contentscrapper.RunContentScapper(url)
		if error != nil {
			return
		}
		data = append(data, newsdata...)
	}

	fmt.Println("Total News Collected: ", len(data))
}
