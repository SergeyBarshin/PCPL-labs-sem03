package fetcher

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// получаем html
func fetchHTML(url string) (*html.Node, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching URL: %v", err)
    }
    defer resp.Body.Close()

    // Парсим HTML
    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error parsing HTML: %v", err)
    }

    return doc, nil
}

// Рекурсивно обходим дерево
func extractLinks(n *html.Node, baseURL string) []string {
    var links []string
    if n.Type == html.ElementNode && n.Data == "a" {
        // Ищем тэг <a> и извлекаем атрибут href
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                href := attr.Val
                // Если ссылка относительная, то не идем туда
                if strings.HasPrefix(href, "http") {
                    links = append(links, href)
                }
                break
            }
        }
    }

    // Обходим дочерние узлы
    for child := n.FirstChild; child != nil; child = child.NextSibling {
        links = append(links, extractLinks(child, baseURL)...)
    }

    return links
}


type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}


func Fetch(url string) (string, []string, error) {
	doc, err := fetchHTML(url)

	if err != nil {
		//log.Printf("Error: %v", err)
		return "", nil, err
	}

	links := extractLinks(doc, url)

	return url, links, err
}
 