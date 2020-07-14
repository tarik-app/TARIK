package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tarik-app/TARIK/touristmedia"
)

type LatitLong struct {
	Latit float64 `json:"latit"`
	Longi float64 `json:"longi"`
}

func Location(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var coor LatitLong
	_ = json.NewDecoder(r.Body).Decode(&coor)
	//
	mediasummeryresult := touristmedia.GetNearbyTouristAttraction(coor.Latit, coor.Longi)
	fmt.Println(mediasummeryresult)
	json.NewEncoder(w).Encode(coor)
}

func main() {
	r := mux.NewRouter()
	port := ":" + os.Getenv("PORT")
	fmt.Println(port)
	fmt.Println("Go WebSocket!")

	r.HandleFunc("/", Location)

	log.Fatal(http.ListenAndServe(port, r))
	// ":8000"
}
