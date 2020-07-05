package touristmedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type WikiArticleTitles struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue      struct {
		Sroffset int    `json:"sroffset"`
		Continue string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Searchinfo struct {
			Totalhits int `json:"totalhits"`
		} `json:"searchinfo"`
		Search []struct {
			Ns     int    `json:"ns"`
			Title  string `json:"title"`
			Pageid int    `json:"pageid"`
		} `json:"search"`
	} `json:"query"`
}

func WikiRequest(query string) string {

	url, err := url.Parse("http://en.wikipedia.org/w/api.php")
	if err != nil {
		log.Fatal(err)
	}
	url.Scheme = "https"
	q := url.Query()
	q.Add("list", "search")
	q.Add("srprop", "")
	q.Add("srsearch", query)
	q.Add("format", "json")
	q.Add("action", "query")

	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)

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

	var titles WikiArticleTitles
	err = json.Unmarshal(bodyBytes, &titles)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(titles.Query.Search[0].Title)

	return titles.Query.Search[0].Title

}
