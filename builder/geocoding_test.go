package builder

import (
	"log"
	"testing"
)

func TestGeocodingUrl(t *testing.T) {
	expected := "http://aa.bb/6.2/search.json?app_id=xx&app_code=yy&gen=9&searchtext=sdf+dfdf+qwe"
	url := NewGeocodingService().SetHost("http://aa.bb").SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestGeocodingUrlWithoutHost(t *testing.T) {
	expected := "https://geocoder.api.here.com/6.2/search.json?app_id=xx&app_code=yy&gen=9&searchtext=sdf+dfdf+qwe"
	url := NewGeocodingService().SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestGeocodingUrlWithoutQuery(t *testing.T) {
	expected := "https://geocoder.api.here.com/6.2/search.json?app_id=xx&app_code=yy&gen=9"
	url := NewGeocodingService().SetAppID("xx").SetAppToken("yy").Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}
