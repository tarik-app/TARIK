package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type NearbyTourSites struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Results          []struct {
		BusinessStatus string `json:"business_status"`
		Geometry       struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Viewport struct {
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
		Icon         string `json:"icon"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		OpeningHours struct {
			OpenNow bool `json:"open_now"`
		} `json:"opening_hours,omitempty"`
		Photos []struct {
			Height           int      `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos"`
		PlaceID  string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		Rating           int      `json:"rating"`
		Reference        string   `json:"reference"`
		Scope            string   `json:"scope"`
		Types            []string `json:"types"`
		UserRatingsTotal int      `json:"user_ratings_total"`
		Vicinity         string   `json:"vicinity"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetNearbyTouristAttraction(lat, long float64) (*http.Response, error) {
	// making API call and returns http response

	fmt.Println("coordinate in get nearby tourist attraction :{ ")
	fmt.Println(lat)
	fmt.Println(long)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%f,%f&radius=500&type=tourist_attraction&keyword=cruise&key=AIzaSyBV8iWuM-TmtoQwN91nBigfreJvys4tTiY", lat, long)
	APIURL := url
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
	println(bodyString)

	var nearby NearbyTourSites
	// storing it in struct to pass it to GetNearbyTouristAttraction function
	err = json.Unmarshal(bodyBytes, &nearby)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("result name >>>")
	// fmt.Println(nearby.Results)
	// fmt.Println(reflect.TypeOf(nearby.Results))
	fmt.Println("---------------------------")
	for i := 0; i < len(nearby.Results); i++ {
		fmt.Println(nearby.Results[i].Name)
		fmt.Println("\n")
	}
	return resp, nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type LatitLong struct {
	Latit float64 `json:"latit"`
	Longi float64 `json:"longi"`
}

func reader(conn *websocket.Conn) {
	// listening for incoming messages
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// getting long and latit as json data from the frontend(index.html file)

		var coor LatitLong
		// storing it in struct to pass it to GetNearbyTouristAttraction function
		err = json.Unmarshal([]byte(p), &coor)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("coordinates")
		fmt.Println("latitude", coor.Latit)
		fmt.Println("longtiude", coor.Longi)

		resp, err := GetNearbyTouristAttraction(coor.Latit, coor.Longi)
		if err != nil {
			fmt.Println("error when finding NearbyTouristAttraction")
		}

		fmt.Println("response after calling GetNearbyTouristAttraction: ")
		fmt.Println(resp)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade an incoming connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Successfully connected.....")
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go WebSocket!")
	// r := mux.NewRouter()
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
