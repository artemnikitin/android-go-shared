package hereapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/artemnikitin/android-go-shared/builder"
)

type geocodingResponse struct {
	Response struct {
		MetaInfo struct {
			Timestamp string `json:"Timestamp"`
		} `json:"MetaInfo"`
		View []struct {
			Result []struct {
				Location struct {
					Address struct {
						AdditionalData []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"AdditionalData"`
						City        string `json:"City"`
						Country     string `json:"Country"`
						County      string `json:"County"`
						District    string `json:"District"`
						HouseNumber string `json:"HouseNumber"`
						Label       string `json:"Label"`
						PostalCode  string `json:"PostalCode"`
						State       string `json:"State"`
						Street      string `json:"Street"`
					} `json:"Address"`
					DisplayPosition struct {
						Latitude  float64 `json:"Latitude"`
						Longitude float64 `json:"Longitude"`
					} `json:"DisplayPosition"`
					LocationID   string `json:"LocationId"`
					LocationType string `json:"LocationType"`
					MapView      struct {
						BottomRight struct {
							Latitude  float64 `json:"Latitude"`
							Longitude float64 `json:"Longitude"`
						} `json:"BottomRight"`
						TopLeft struct {
							Latitude  float64 `json:"Latitude"`
							Longitude float64 `json:"Longitude"`
						} `json:"TopLeft"`
					} `json:"MapView"`
					NavigationPosition []struct {
						Latitude  float64 `json:"Latitude"`
						Longitude float64 `json:"Longitude"`
					} `json:"NavigationPosition"`
				} `json:"Location"`
				MatchLevel   string `json:"MatchLevel"`
				MatchQuality struct {
					City        float64   `json:"City"`
					HouseNumber float64   `json:"HouseNumber"`
					Street      []float64 `json:"Street"`
				} `json:"MatchQuality"`
				MatchType string  `json:"MatchType"`
				Relevance float64 `json:"Relevance"`
			} `json:"Result"`
			ViewID int    `json:"ViewId"`
			Type   string `json:"_type"`
		} `json:"View"`
	} `json:"Response"`
}

// GetCoordinates implements HERE Geocoding API for convert text address to GPS coordinates
func GetCoordinates(appID, appToken, searchText string) string {
	builder := builder.NewGeocodingService().SetHost("https://geocoder.cit.api.here.com")
	url := builder.SetAppID(appID).SetAppToken(appToken).SetSearchPhrase(searchText).Build()
	resp := sendRequest(url)
	defer closeAfter(resp)
	if resp.StatusCode != 200 {
		return ""
	}
	bytes := getBody(resp)
	lat, lon := getCoordinatesFromJSON(bytes)
	return createStringFromCoordinates(lat, lon)
}

// GetPicture returns map tile for specific set of GPS coordinates
func GetPicture(appID, appToken string, lat, lon float64, h, w, dpi int) []byte {
	builder := builder.NewMapTileService().SetHost("https://image.maps.cit.api.here.com").SetAppID(appID).SetAppToken(appToken)
	url := builder.SetLatitude(lat).SetLongitude(lon).SetWidth(w).SetHeight(h).SetDpi(dpi).Build()
	resp := sendRequest(url)
	defer closeAfter(resp)
	if resp.StatusCode != 200 {
		return make([]byte, 0)
	}
	return getBody(resp)
}

func sendRequest(data string) *http.Response {
	resp, err := http.Get(data)
	if err != nil {
		log.Println("Can't execute HTTP request ...")
		log.Println(err)
		return resp
	}
	return resp
}

func getBody(resp *http.Response) []byte {
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Can't get a body of HTTP response ...")
		log.Println(err)
	}
	return response
}

func closeAfter(resp *http.Response) {
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
}

func getCoordinatesFromJSON(response []byte) (float64, float64) {
	var geocode = &geocodingResponse{}
	var lat, lon float64
	err := json.Unmarshal(response, geocode)
	if err != nil {
		log.Println("Can't parse JSON ...")
		log.Println(err)
		return lat, lon
	}
	lat = geocode.Response.View[0].Result[0].Location.DisplayPosition.Latitude
	lon = geocode.Response.View[0].Result[0].Location.DisplayPosition.Longitude
	return lat, lon
}

func createStringFromCoordinates(lat, lon float64) string {
	var bytes bytes.Buffer
	bytes.WriteString(strconv.FormatFloat(lat, 'f', -1, 64))
	bytes.WriteString("||")
	bytes.WriteString(strconv.FormatFloat(lon, 'f', -1, 64))
	return bytes.String()
}
