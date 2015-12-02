package hereapi

import (
	"log"
	"os"
	"testing"
)

var (
	appID    = os.Getenv("HERE_ANDROID_GO_ID")
	appToken = os.Getenv("HERE_ANDROID_GO_TOKEN")
)

func TestGetMapPicturePositive(t *testing.T) {
	image := GetPicture(appID, appToken, 52.5308599, 13.38469, 320, 240, 24)
	if image[0] != 0xFF || image[1] != 0xD8 {
		t.Error("Should return an image in response")
	}
}

func TestGetMapPictureUnauthorized(t *testing.T) {
	image := GetPicture("zxz", "w", 52.5308599, 13.38469, 32, 24, 24)
	if len(image) != 0 {
		t.Error("Should return an empty byte array with incorrect authorization data")
	}
}

func TestParseJson(t *testing.T) {
	realJSON := []byte(`{"Response":{"MetaInfo":{"Timestamp":"2015-11-24T16:24:45.952+0000"},"View":[{"_type":"SearchResultsViewType","ViewId":0,"Result":[{"Relevance":1.0,"MatchLevel":"houseNumber","MatchQuality":{"City":1.0,"Street":[1.0],"HouseNumber":1.0},"MatchType":"pointAddress","Location":{"LocationId":"NT_5v-U9-t5AZa5VsRdV06lLB_xEjN","LocationType":"address","DisplayPosition":{"Latitude":52.5308599,"Longitude":13.38469},"NavigationPosition":[{"Latitude":52.53098,"Longitude":13.38458}],"MapView":{"TopLeft":{"Latitude":52.5319841,"Longitude":13.3828421},"BottomRight":{"Latitude":52.5297357,"Longitude":13.3865379}},"Address":{"Label":"Invalidenstraße 116, 10115 Berlin, Deutschland","Country":"DEU","State":"Berlin","County":"Berlin","City":"Berlin","District":"Mitte","Street":"Invalidenstraße","HouseNumber":"116","PostalCode":"10115","AdditionalData":[{"value":"Deutschland","key":"CountryName"},{"value":"Berlin","key":"StateName"},{"value":"Berlin","key":"CountyName"}]}}}]}]}}`)
	lat, lon := getCoordinatesFromJSON(realJSON)
	if lat != 52.5308599 || lon != 13.38469 {
		log.Fatal(lat)
		log.Fatal(lon)
		t.Error("JSON was incorrectly parsed")
	}
}

func TestGetCoordinatesRequest(t *testing.T) {
	coordinates := GetCoordinates(appID, appToken, "Berlin Invalidenstrasse 116")
	if coordinates != "52.5308599||13.38469" {
		log.Fatal(coordinates)
		t.Error("Incorrect coordinates")
	}
}

func TestGetCoordinatesUnauthorized(t *testing.T) {
	coordinates := GetCoordinates("xx", "yy", "Berlin Invalidenstrasse 116")
	if coordinates != "" {
		log.Fatal(coordinates)
		t.Error("Should return 0 for unathorized request")
	}
}

func TestCreateStringFromCoordinates(t *testing.T) {
	str := createStringFromCoordinates(23.434343, 27.78990890976)
	if str != "23.434343||27.78990890976" {
		log.Fatal(str)
		t.Error("Incorrect creation of string from coordinates")
	}
}
