package parsehtml

import (
	newsdata "google-news-scrapper/types"

	"golang.org/x/net/html"
)

func ParseBody(n *html.Node, list_of_news *[]newsdata.Data) {
	if n.Type == html.ElementNode && n.Data == "article" {
		newsDataCollector := newsdata.Data{}
		ParseArticle(n, &newsDataCollector)
		*list_of_news = append(*list_of_news, newsDataCollector)
	}

	// traverse the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseBody(c, list_of_news)
	}
}

func ParseArticle(n *html.Node, newsDataCollector *newsdata.Data) {
	switch n.Data {
	case "img":
		// Thumbnail
		isThumbnail := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "tvs3Id QwxBBf" {
				isThumbnail = true
			}
			if isThumbnail && a.Key == "src" {
				newsDataCollector.Thumbnail = a.Val
			}
		}
	case "a":
		// title
		isTitle := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "DY5T1d RZIKme" {
				isTitle = true
			}
			if isTitle && n.FirstChild.Type == html.TextNode {
				newsDataCollector.Title = n.FirstChild.Data
			}
		}
	case "h4":
		// source
		isUrl := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "ipQwMb ekueJc RD0gLb" {
				isUrl = true
			}
		}

		if isUrl {
			anchor := n.FirstChild
			for _, a := range anchor.Attr {
				if a.Key == "href" {
					newsDataCollector.Url = a.Val[1:]
				}
			}
		}
	case "div":
		// source
		isSource := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "wsLqz RD0gLb" {
				isSource = true
			}
		}
		if isSource {
			img := n.FirstChild
			if img.Data == "img" {
				for _, a := range img.Attr {
					if a.Key == "src" {
						newsDataCollector.Source.Logo = a.Val
					}
				}
			}

			div := n.FirstChild.NextSibling
			anchor := div.FirstChild
			if div.Data == "div" && anchor.Data == "a" && anchor.FirstChild.Type == html.TextNode {
				newsDataCollector.Source.Name = anchor.FirstChild.Data
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseArticle(c, newsDataCollector)
	}
}

func ParseArticle2(n *html.Node, newsDataCollector *newsdata.Data) {
	switch n.Data {
	case "img":
		// Thumbnail
		isThumbnail := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "tvs3Id QwxBBf" {
				isThumbnail = true
			}
			if isThumbnail && a.Key == "src" {
				newsDataCollector.Thumbnail = a.Val
			}
		}
	case "a":
		// title
		isTitle := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "DY5T1d RZIKme" {
				isTitle = true
			}
			if isTitle && n.FirstChild.Type == html.TextNode {
				newsDataCollector.Title = n.FirstChild.Data
			}
		}
	case "h4":
		// source
		isUrl := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "ipQwMb ekueJc RD0gLb" {
				isUrl = true
			}
		}

		if isUrl {
			anchor := n.FirstChild
			for _, a := range anchor.Attr {
				if a.Key == "href" {
					newsDataCollector.Url = a.Val[1:]
				}
			}
		}
	case "div":
		// source
		isSource := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "wsLqz RD0gLb" {
				isSource = true
			}
		}
		if isSource {
			img := n.FirstChild
			if img.Data == "img" {
				for _, a := range img.Attr {
					if a.Key == "src" {
						newsDataCollector.Source.Logo = a.Val
					}
				}
			}

			div := n.FirstChild.NextSibling
			anchor := div.FirstChild
			if div.Data == "div" && anchor.Data == "a" && anchor.FirstChild.Type == html.TextNode {
				newsDataCollector.Source.Name = anchor.FirstChild.Data
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseArticle(c, newsDataCollector)
	}
}
