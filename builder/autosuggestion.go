package builder

import (
	"bytes"
	"strconv"
	"strings"
)

// AutosuggestionService interface for builder
type AutosuggestionService interface {
	SetHost(string) AutosuggestionService
	SetAppID(string) AutosuggestionService
	SetAppToken(string) AutosuggestionService
	SetSearchPhrase(string) AutosuggestionService
	SetMaxResults(int) AutosuggestionService
	SetLatitude(float64) AutosuggestionService
	SetLongitude(float64) AutosuggestionService
	Build() string
}

type autosuggestionService struct {
	host         string
	appID        string
	appToken     string
	searchString string
	maxResults   int
	latitude     float64
	longitude    float64
}

// NewAutosuggestionService return new builder
func NewAutosuggestionService() AutosuggestionService {
	return &autosuggestionService{
		host:       "https://autocomplete.geocoder.api.here.com",
		maxResults: 10,
		latitude:   0,
		longitude:  0,
	}
}

func (as *autosuggestionService) SetHost(host string) AutosuggestionService {
	as.host = host
	return as
}

func (as *autosuggestionService) SetAppID(id string) AutosuggestionService {
	as.appID = id
	return as
}

func (as *autosuggestionService) SetAppToken(token string) AutosuggestionService {
	as.appToken = token
	return as
}

func (as *autosuggestionService) SetSearchPhrase(text string) AutosuggestionService {
	as.searchString = strings.Replace(text, " ", "+", -1)
	return as
}

func (as *autosuggestionService) SetMaxResults(res int) AutosuggestionService {
	as.maxResults = res
	return as
}

func (as *autosuggestionService) SetLatitude(coordinate float64) AutosuggestionService {
	as.latitude = coordinate
	return as
}

func (as *autosuggestionService) SetLongitude(coordinate float64) AutosuggestionService {
	as.longitude = coordinate
	return as
}

func (as *autosuggestionService) Build() string {
	var buffer bytes.Buffer
	buffer.WriteString(as.host)
	buffer.WriteString("/6.2/suggest.json?app_id=")
	buffer.WriteString(as.appID)
	buffer.WriteString("&app_code=")
	buffer.WriteString(as.appToken)
	buffer.WriteString("&query=")
	buffer.WriteString(as.searchString)
	buffer.WriteString("&maxresults=")
	buffer.WriteString(strconv.Itoa(as.maxResults))
	if as.latitude != 0 && as.longitude != 0 {
		buffer.WriteString("&prox=")
		buffer.WriteString(strconv.FormatFloat(as.latitude, 'f', -1, 64))
		buffer.WriteString(",")
		buffer.WriteString(strconv.FormatFloat(as.longitude, 'f', -1, 64))
	}
	return buffer.String()
}
