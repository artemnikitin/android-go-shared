package hereapi

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"antonholmquist/jason"
	"github.com/artemnikitin/android-go-shared/builder"
)

func GetCoordinates(appID, appToken, searchText string) string {
	builder := builder.NewGeocodingService()
	builder = builder.SetHost("http://geocoder.cit.api.here.com").SetAppID(appID).SetAppToken(appToken)
	url := builder.SetSearchPhrase(searchText).Build()
	var result string
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Can't execute HTTP request ...")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			log.Fatal("Can't get a JSON response ...")
		}
		lat, lon := getCoordinatesFromJSON(bytes)
		result = createStringFromCoordinates(lat, lon)
	}
	return result
}

func GetPicture(appID, appToken string, lat, lon float64, h, w, dpi int) []byte {
	builder := builder.NewMapTileService()
	builder = builder.SetHost("http://image.maps.cit.api.here.com").SetAppID(appID).SetAppToken(appToken)
	url := builder.SetLatitude(lat).SetLongitude(lon).SetWidth(w).SetHeight(h).SetDpi(dpi).Build()
	var response []byte
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Can't execute HTTP request ...")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		response, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			log.Fatal("Can't get a body of HTTP response ...")
		}
	}
	return response
}

func getCoordinatesFromJSON(response []byte) (float64, float64) {
	var lat, lon float64
	js, err := jason.NewObjectFromBytes(response)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Can't parse JSON ...")
	}
	resp, _ := js.GetObject("Response")
	view, _ := resp.GetObjectArray("View")
	result, _ := view[0].GetObjectArray("Result")
	loc, _ := result[0].GetObject("Location")
	display, _ := loc.GetObject("DisplayPosition")
	lat, _ = display.GetFloat64("Latitude")
	lon, _ = display.GetFloat64("Longitude")
	return lat, lon
}

func createStringFromCoordinates(lat, lon float64) string {
	var bytes bytes.Buffer
	bytes.WriteString(strconv.FormatFloat(lat, 'f', -1, 64))
	bytes.WriteString("||")
	bytes.WriteString(strconv.FormatFloat(lon, 'f', -1, 64))
	return bytes.String()
}
