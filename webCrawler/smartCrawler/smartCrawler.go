package main

import (
	"crawler/URLCache"
	"crawler/fetcher"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

/*
Многопочный кравлер, с кешем
*/


var cache = URLCache.URLCache{Cache: make(map[string]bool)}


func Crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()


	if depth <= 0 || cache.Visit(url){
		return
	}
	
	// находим на странице ссылки
	_, urls, err := fetcher.Fetch(url)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

    log.Printf("%d: %s\n", depth, url)


	for _, u := range urls {
		if(depth > 1){
			wg.Add(1)
			go Crawl(u, depth-1, wg)
		}
	}
}


func main() {
	godotenv.Load()
    depth, _ := strconv.Atoi(os.Getenv("DEPTH"))
    startURl := os.Getenv("START_URl")


	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl(startURl, depth, &wg)
	wg.Wait()


	fmt.Printf("\n ПОСЕЩЕНО %d страниц\n", len(cache.Cache))
}

/*
и того: в мейн вызываем первый crawl, он "парсит" страницу, вызываем crawl для всех его потомков, 
*/

