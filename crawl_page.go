package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.getPagesLength() >= cfg.maxPages {
		return
	}

	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("couldn't parse the current url: %v", err)
		return
	}

	if cfg.baseURL.Hostname() != currURL.Hostname() {
		return
	}

	normalizedCurrURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("couldn't normalize the current url: %v", err)
		return
	}

	cfg.mu.Lock()
	if count, visited := cfg.pages[normalizedCurrURL]; visited {
		cfg.pages[normalizedCurrURL] = count + 1
		cfg.mu.Unlock()
		return
	}

	cfg.pages[normalizedCurrURL] = 1
	cfg.mu.Unlock()

	fmt.Printf("currently crawling %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("couldn't get html content: %v", err)
		return
	}

	fmt.Printf("retrieved html content from %s\n", rawCurrentURL)

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		log.Printf("couldn't extract urls from html: %v", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
