package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func GetMediaWiki() (*http.Response, error) {
	// making API call and returns http response
	APIURL := "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=Alcatraz%20Island"
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// fmt.Println(reflect.TypeOf(resp))
	return resp, nil

}

func MediaWikiHandler(w http.ResponseWriter, r *http.Request) {
	// receives http response and post to the server side
	resp, err := GetMediaWiki()
	if err != nil {
		fmt.Println("error from GetMediaWiki")
	}

	mediawiki := MediaWiki{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// source https://www.golanglearn.com/how-to-parse-json-data-in-golang/
	var wiki MediaWiki
	decoder := json.NewDecoder(resp.Body).Decode(&wiki)
	fmt.Printf("%+v\n", mediawiki)
	fmt.Println(decoder)
	res, err := json.Marshal(&wiki)
	w.Write(res)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("wiki.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	resp, err := GetMediaWiki()
	if err != nil {
		fmt.Println("error from GetMediaWiki")
	}

	mediawiki := MediaWiki{}

	var wiki MediaWiki
	decoder := json.NewDecoder(resp.Body).Decode(&wiki)
	fmt.Printf("%+v\n", mediawiki)
	fmt.Println(decoder)

	t.Execute(w, &wiki)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/wiki", MediaWikiHandler)
	r.HandleFunc("/template", templateHandler)
	log.Fatal(http.ListenAndServe(":8010", r))
}
