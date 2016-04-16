package builder

import (
	"bytes"
	"strings"
)

// GeocodingService interface for builder
type GeocodingService interface {
	SetHost(string) GeocodingService
	SetAppID(string) GeocodingService
	SetAppToken(string) GeocodingService
	SetSearchPhrase(string) GeocodingService
	Build() string
}

type geocodingService struct {
	host         string
	appID        string
	appToken     string
	searchString string
}

// NewGeocodingService return new builder
func NewGeocodingService() GeocodingService {
	return &geocodingService{host: "https://geocoder.cit.api.here.com"}
}

func (gs *geocodingService) SetHost(host string) GeocodingService {
	gs.host = host
	return gs
}

func (gs *geocodingService) SetAppID(id string) GeocodingService {
	gs.appID = id
	return gs
}

func (gs *geocodingService) SetAppToken(token string) GeocodingService {
	gs.appToken = token
	return gs
}

func (gs *geocodingService) SetSearchPhrase(text string) GeocodingService {
	gs.searchString = strings.Replace(text, " ", "+", -1)
	return gs
}

func (gs *geocodingService) Build() string {
	var buffer bytes.Buffer
	buffer.WriteString(gs.host)
	buffer.WriteString("/6.2/search.json?app_id=")
	buffer.WriteString(gs.appID)
	buffer.WriteString("&app_code=")
	buffer.WriteString(gs.appToken)
	buffer.WriteString("&searchtext=")
	buffer.WriteString(gs.searchString)
	buffer.WriteString("&gen=9")
	return buffer.String()
}
