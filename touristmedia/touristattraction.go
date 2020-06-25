package touristmedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetNearbyTouristAttraction(lat, long float64) {
	// making API call and returns http response
	APIURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%f,%f&radius=1000&type=tourist_attraction&keyword=cruise&key=AIzaSyBV8iWuM-TmtoQwN91nBigfreJvys4tTiY", lat, long)
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

	// var wiki MediaWiki
	for i := 0; i < len(nearby.Results); i++ {
		// "Painted ladies"
		//
		mediaResp, mediaErr := GetMediaWiki(nearby.Results[i].Name)
		if mediaErr != nil {
			fmt.Println("Error from Media......")
		}

		mediaBodyBytes, _ := ioutil.ReadAll(mediaResp.Body)
		mediaBodyString := string(mediaBodyBytes)
		fmt.Println("***********************************************************************")
		print(mediaBodyString)

		var wiki MediaWiki
		// storing it in struct to pass it to GetNearbyTouristAttraction function
		err = json.Unmarshal(mediaBodyBytes, &wiki)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(wiki.Query.Pages.Num64107.Extract)

	}
}
