package touristmedia

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type MediaWiki struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages struct {
			Num64107 struct {
				Pageid  int    `json:"pageid"`
				Ns      int    `json:"ns"`
				Title   string `json:"title"`
				Extract string `json:"extract"`
			} `json:"64107"`
		} `json:"pages"`
	} `json:"query"`
}

func GetMediaWiki(site string) string {
	fmt.Println(site)

	query := WikiRequest(site)
	safeQuery := url.QueryEscape(query)
	APIURL := fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s", safeQuery)

	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	// println(bodyString)
	return bodyString

}
