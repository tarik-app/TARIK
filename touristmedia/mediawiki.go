package touristmedia

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
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

// china town SF
func GetMediaWiki(site string) (*http.Response, error) {
	fmt.Println(site)

	// cmd := exec.Command("python3", "-c", fmt.Sprintf("import wikisearch; print(wikisearch.wiki_search(\"%s\"))", site))
	// cmd := exec.Command("python", "-c", fmt.Sprintf("import wikisearch; print(wikisearch.wiki_search(\"%s\"))", site))
	// fmt.Println(cmd.Args)
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// s := strings.Split(string(out), ",")
	// fmt.Println(s)
	// fmt.Println(s[0])
	cmd := exec.Command("wikisearch.py")
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
		// return
	}
	fmt.Println(string(out))

	safeQuery := url.QueryEscape("painted ladies")
	APIURL := fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s", safeQuery)

	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}
	// Chinatown, San Francisco
	// The Painted ladies

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp, nil

}
