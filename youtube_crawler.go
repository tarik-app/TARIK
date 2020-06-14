package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}

	netCilent = &http.Client{
		Transport: transport,
	}

	queue     = make(chan string)
	isVisited = make(map[string]bool)
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("missing url, e.g  go run history_scraper.go https://youtube.com/jsfunc")
		os.Exit(1)
	}

	baseURL := args[0]
	go func() {
		queue <- baseURL
	}()

	for href := range queue {
		if !isVisited[href] && isSameDomain(href, baseURL) {
			crawlURL(href)
		}
	}
}

func crawlURL(href string) {
	isVisited[href] = true
	fmt.Printf("Crawling url --> %v \n", href)
	response, err := netCilent.Get(href)
	defer response.Body.Close()
	checkErr(err)

	links, err := extractlinks.All(response.Body)
	checkErr(err)
	for _, link := range links {
		absoluteURL := toFixedURL(link.Href, href)
		// concurrent
		go func() {
			// recive
			queue <- absoluteURL
			// crawlURL()
		}()

	}

}

func isSameDomain(href, baseURL string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}

	fmt.Println("URI host", uri.Host)

	parentUri, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host {
		return false
	}

	return true
}

func toFixedURL(href, baseURL string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}

	fmt.Println("URI host", uri.Host)

	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	fmt.Println("base host", base.Host)

	toFixedURI := base.ResolveReference(uri)

	return toFixedURI.String()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		// exist from terminal
		os.Exit(1)
	}
}
