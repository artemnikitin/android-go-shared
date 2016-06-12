package builder

import (
	"bytes"
	"strconv"
)

var tileParams map[string]string

// MapTileService interface for builder
type MapTileService interface {
	SetHost(string) MapTileService
	SetAppID(string) MapTileService
	SetAppToken(string) MapTileService
	SetLatitude(float64) MapTileService
	SetLongitude(float64) MapTileService
	SetHeight(int) MapTileService
	SetWidth(int) MapTileService
	SetDpi(int) MapTileService
	Build() string
}

type mapService struct {
	host      string
	appID     string
	appToken  string
	latitude  float64
	longitude float64
}

// NewMapTileService return new builder
func NewMapTileService() MapTileService {
	tileParams = make(map[string]string)
	return &mapService{host: "https://image.maps.api.here.com"}
}

func (ms *mapService) SetHost(host string) MapTileService {
	ms.host = host
	return ms
}

func (ms *mapService) SetAppID(id string) MapTileService {
	ms.appID = id
	return ms
}

func (ms *mapService) SetAppToken(token string) MapTileService {
	ms.appToken = token
	return ms
}

func (ms *mapService) SetLatitude(coordinate float64) MapTileService {
	ms.latitude = coordinate
	return ms
}

func (ms *mapService) SetLongitude(coordinate float64) MapTileService {
	ms.longitude = coordinate
	return ms
}

func (ms *mapService) SetHeight(value int) MapTileService {
	tileParams["h"] = strconv.Itoa(value)
	return ms
}

func (ms *mapService) SetWidth(value int) MapTileService {
	tileParams["w"] = strconv.Itoa(value)
	return ms
}

func (ms *mapService) SetDpi(value int) MapTileService {
	tileParams["ppi"] = strconv.Itoa(value)
	return ms
}

func (ms *mapService) Build() string {
	var buffer bytes.Buffer
	buffer.WriteString(ms.host)
	buffer.WriteString("/mia/1.6/mapview?app_id=")
	buffer.WriteString(ms.appID)
	buffer.WriteString("&app_code=")
	buffer.WriteString(ms.appToken)
	buffer.WriteString("&c=")
	buffer.WriteString(strconv.FormatFloat(ms.latitude, 'f', -1, 64))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatFloat(ms.longitude, 'f', -1, 64))
	buffer.WriteString("&z=18&u=10")
	for k, v := range tileParams {
		buffer.WriteString("&")
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
	}
	return buffer.String()
}
