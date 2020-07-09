package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

type LatitLong struct {
	Latit float64 `json:"latit"`
	Longi float64 `json:"longi"`
}

// func reader(conn *websocket.Conn) {
// 	// listening for incoming messages
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		// getting long and latit as json data
// 		var coor LatitLong
// 		// storing it in struct
// 		err = json.Unmarshal([]byte(p), &coor)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		mediasummeryresult := touristmedia.GetNearbyTouristAttraction(coor.Latit, coor.Longi)
// 		fmt.Println(mediasummeryresult)

// 		if err := conn.WriteJSON(mediasummeryresult); err != nil {
// 			log.Println(err)
// 			break
// 		}

// 	}
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	// upgrade an incoming connection
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

// 	fmt.Println(r.Host)

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("Client Successfully connected.....")
// 	reader(ws)
// }

func Location(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var coor LatitLong
	_ = json.NewDecoder(r.Body).Decode(&coor)
	fmt.Println(coor, r.Body)
	json.NewEncoder(w).Encode(coor)
}

func main() {
	fmt.Println("Go WebSocket!")
	http.HandleFunc("/location", Location)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
