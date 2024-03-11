package parseurl

import (
	"golang.org/x/net/html"
)

func ParseBody(n *html.Node, list_of_url *[]string) {
	if n.Type == html.ElementNode {
		ParseURL(n, list_of_url)
	}

	// traverse the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseBody(c, list_of_url)
	}
}

func ParseURL(n *html.Node, list_of_url *[]string) {
	switch n.Data {
	case "a":
		isURL := false
		for _, a := range n.Attr {
			if a.Key == "class" && (a.Val == "jKHa4e" || a.Val == "bwOy4d VfPpkd-ksKsZd-XxIAqe") {
				isURL = true
			}
			if isURL && a.Key == "href" {
				*list_of_url = append(*list_of_url, "https://news.google.com"+a.Val[1:])
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseURL(c, list_of_url)
	}
}
