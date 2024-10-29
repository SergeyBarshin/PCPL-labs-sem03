package main

import (
	"crawler/URLCache"
	"crawler/fetcher"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/*
однопоточный кравлер с кешем
*/


var cache = URLCache.URLCache{Cache: make(map[string]bool)}

func Crawl(url string, depth int) {

	if depth <= 0 || cache.Visit(url) {
		return
	}

	_, urls, err := fetcher.Fetch(url)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

    log.Printf("%d: %s\n", depth, url)

	for _, u := range urls {
		if(depth > -1) {
			Crawl(u, depth-1)
		}
		
	}

	return
}

func main() {
	godotenv.Load()
    depth, _ := strconv.Atoi(os.Getenv("DEPTH"))
    startURl := os.Getenv("START_URl")


	Crawl(startURl, depth)
	fmt.Printf("\n ПОСЕЩЕНО %d страниц\n", len(cache.Cache))

}
