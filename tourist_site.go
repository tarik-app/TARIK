package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTouristAttraction() (*http.Response, error) {
	// making API call and returns http response
	APIURL := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=san+francisco+city+point+of+interest&language=en&key=AIzaSyBV8iWuM-TmtoQwN91nBigfreJvys4tTiY"
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

// json to go
// map the json into a struct to map objects
type TouristAttractionSites struct {
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

func TouristAttractionHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := GetTouristAttraction()
	if err != nil {
		fmt.Println("error from GetMediaWiki")
	}

	toursit_sites := TouristAttractionSites{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// source https://www.golanglearn.com/how-to-parse-json-data-in-golang/
	var sites TouristAttractionSites
	decoder := json.NewDecoder(resp.Body).Decode(&sites)
	fmt.Printf("%+v\n", toursit_sites)
	fmt.Println(decoder)
	res, err := json.Marshal(&sites)
	w.Write(res)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tourist-site", TouristAttractionHandler)
	log.Fatal(http.ListenAndServe(":8100", r))
}
