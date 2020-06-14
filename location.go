package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// json to go
// map the json into a struct to map objects
type GoogleGeoLocation struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID  string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		Types []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

func GeoLocationHandler(w http.ResponseWriter, r *http.Request) {

	APIURL := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=san+francisco+city+point+of+interest&language=en&key=AIzaSyBV8iWuM-TmtoQwN91nBigfreJvys4tTiY"
	// "https://maps.googleapis.com/maps/api/geocode/json?address=Chinatown,+CA&key=AIzaSyBV8iWuM-TmtoQwN91nBigfreJvys4tTiY"
	// "https://www.googleapis.com/customsearch/v1?key=AIzaSyBN6P9Qqj7BVzYBdJZCML3phYUPAtg-ZUM&cx=017576662512468239146:omuauf_lfve&q=cars&callback=hndlr"
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}
	// fmt.Println(req)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	geolocation := GoogleGeoLocation{}

	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// source https://www.golanglearn.com/how-to-parse-json-data-in-golang/
	var geo GoogleGeoLocation
	decoder := json.NewDecoder(resp.Body).Decode(&geo)
	fmt.Printf("%+v\n", geolocation)
	fmt.Println(decoder)
	res, err := json.Marshal(&geo)
	w.Write(res)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/geoloc", GeoLocationHandler)
	log.Fatal(http.ListenAndServe(":8100", r))
}
