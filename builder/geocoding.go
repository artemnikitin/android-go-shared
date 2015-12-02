package builder

import (
	"bytes"
	"strings"
)

type GeocodingService interface {
	SetHost(string) GeocodingService
	SetAppId(string) GeocodingService
	SetAppToken(string) GeocodingService
	SetSearchPhrase(string) GeocodingService
	Build() string
}

type geocodingService struct {
	host         string
	appId        string
	appToken     string
	searchString string
}

func NewGeocodingService() GeocodingService {
	return &geocodingService{}
}

func (gs *geocodingService) SetHost(host string) GeocodingService {
	gs.host = host
	return gs
}

func (gs *geocodingService) SetAppId(id string) GeocodingService {
	gs.appId = id
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
	buffer.WriteString("/6.2/geocode.json?app_id=")
	buffer.WriteString(gs.appId)
	buffer.WriteString("&app_code=")
	buffer.WriteString(gs.appToken)
	buffer.WriteString("&searchtext=")
	buffer.WriteString(gs.searchString)
	buffer.WriteString("&gen=9")
	return buffer.String()
}
