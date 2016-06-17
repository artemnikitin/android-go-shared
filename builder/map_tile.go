package builder

import (
	"bytes"
	"strconv"
)

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
	SetParameters(map[string]string) MapTileService
	Build() string
	NewBuild() string
}

type mapService struct {
	host      string
	appID     string
	appToken  string
	latitude  float64
	longitude float64
	params    map[string]string
}

// NewMapTileService return new builder
func NewMapTileService() MapTileService {
	return &mapService{
		host:   "https://image.maps.api.here.com",
		params: make(map[string]string),
	}
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
	ms.params["h"] = strconv.Itoa(value)
	return ms
}

func (ms *mapService) SetWidth(value int) MapTileService {
	ms.params["w"] = strconv.Itoa(value)
	return ms
}

func (ms *mapService) SetDpi(value int) MapTileService {
	ms.params["ppi"] = strconv.Itoa(value)
	return ms
}

func (ms *mapService) SetParameters(params map[string]string) MapTileService {
	ms.params = params
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
	for k, v := range ms.params {
		buffer.WriteString("&")
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
	}
	return buffer.String()
}

func (ms *mapService) NewBuild() string {
	var buffer bytes.Buffer
	buffer.WriteString(ms.host)
	buffer.WriteString("/mia/1.6/mapview?app_id=")
	buffer.WriteString(ms.appID)
	buffer.WriteString("&app_code=")
	buffer.WriteString(ms.appToken)
	for k, v := range ms.params {
		buffer.WriteString("&")
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
	}
	return buffer.String()
}
