package touristmedia

import (
	"fmt"
	"net/http"
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

func GetMediaWiki(site string) (*http.Response, error) {

	APIURL := fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s", site)
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp, nil

}
