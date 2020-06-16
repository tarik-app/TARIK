package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// https://socketloop.com/tutorials/golang-detect-user-location-with-html5-geo-location

func home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Println("hello there!")
	var templates = template.Must(template.New("geolocate").ParseFiles("geolocation.html"))

	err := templates.ExecuteTemplate(w, "geolocation.html", nil)

	if err != nil {
		panic(err)
	}

	prompt := "Detecting your location. Please click 'Allow' button"
	w.Write([]byte(prompt))

}

func location(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lat := vars["lat"]
	long := vars["long"]

	w.Write([]byte(fmt.Sprintf("Lat is %s \n", lat)))

	w.Write([]byte(fmt.Sprintf("Long is %s \n", long)))

	fmt.Printf("Lat is %s and Long is %s \n", lat, long)

	// if you want to get timezone from latitude and longitude
	// checkout http://www.geonames.org/export/web-services.html#timezone

}

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/location/{lat}/{long}", location)

	http.ListenAndServe(":8090", mux)
}
