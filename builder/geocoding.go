package builder

import (
	"bytes"
	"strings"
)

var geoParams map[string]string

// GeocodingService interface for builder
type GeocodingService interface {
	SetHost(string) GeocodingService
	SetAppID(string) GeocodingService
	SetAppToken(string) GeocodingService
	SetSearchPhrase(string) GeocodingService
	Build() string
}

type geocodingService struct {
	host     string
	appID    string
	appToken string
}

// NewGeocodingService return new builder
func NewGeocodingService() GeocodingService {
	geoParams = make(map[string]string)
	return &geocodingService{host: "https://geocoder.api.here.com"}
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
	geoParams["searchtext"] = strings.Replace(text, " ", "+", -1)
	return gs
}

func (gs *geocodingService) Build() string {
	var buffer bytes.Buffer
	buffer.WriteString(gs.host)
	buffer.WriteString("/6.2/search.json?app_id=")
	buffer.WriteString(gs.appID)
	buffer.WriteString("&app_code=")
	buffer.WriteString(gs.appToken)
	buffer.WriteString("&gen=9")
	for k, v := range geoParams {
		buffer.WriteString("&")
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
	}
	return buffer.String()
}
