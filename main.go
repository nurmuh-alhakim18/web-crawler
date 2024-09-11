package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		url := args[0]
		maxConcurrency, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("second argument maxConcurrency: %v", maxConcurrency)
		}

		maxPages, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("second argument maxPages: %v", maxPages)
		}

		fmt.Printf("starting crawl of: %v\n", url)
		cfg, err := configure(url, maxConcurrency, maxPages)
		if err != nil {
			fmt.Printf("Can't configure: %v", err)
		}

		cfg.wg.Add(1)
		go cfg.crawlPage(url)
		cfg.wg.Wait()

		printReport(cfg.pages, url)
	}
}
