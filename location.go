package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

func GetRequest() {
	APIURL := "https://maps.googleapis.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&key=AIzaSyBV8iWuM-TmtoQwN91nBigfreJvys4tTiY"
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
	// fmt.Println(resp)
	geolocation := GoogleGeoLocation{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	// bytes_route, err := json.Marshal(route)
	err2 := json.Unmarshal(body, &geolocation)
	if err2 != nil {
		fmt.Println("error occured during unmarshalling: ", err2)
	}
	// fmt.Println("geolocation after: ", geolocation)
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
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// source https://www.golanglearn.com/how-to-parse-json-data-in-golang/
	var geo GoogleGeoLocation
	decoder := json.NewDecoder(resp.Body).Decode(&geo)
	fmt.Printf("%+v\n", geolocation)
	fmt.Println(decoder)
	res, err := json.Marshal(&geo)
	w.Write(res)
	// jsn, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatal("Error reading the body", err)
	// }
}

type TarikPage struct {
	HistoryOf   string
	Description string
}

func home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Whoa, this place is neat!</h1>")
}

func TarikHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(GetRequest())
	tarik := TarikPage{HistoryOf: "ChinaTown", Description: "For the first Chinatown in the world...."}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, tarik)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// // attention: If you do not call ParseForm method, the following data can not be obtained form
	// fmt.Println(r.Form) // print information on server side.
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	// fmt.Fprintf(w, "Hello astaxie!") // write data to response
	t, _ := template.ParseFiles("geolocator.html")
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Println("method in GET:", r.Method) //get request method
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("method in POST:", r.Method) //get request method
		fmt.Println("Location-Query:", r.Form["location"])
	}
}

func main() {
	GetRequest()
	// bytes_route := []byte(route)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/tarik", TarikHandler)
	r.HandleFunc("/geoloc", GeoLocationHandler)
	r.HandleFunc("/hello", sayhelloName)
	r.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8100", r))
}
