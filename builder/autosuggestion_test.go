package builder

import (
	"log"
	"testing"
)

func TestAutosuggestionUrl(t *testing.T) {
	expected := "http://aa.bb/6.2/suggest.json?app_id=xx&app_code=yy&query=sdf+dfdf+qwe&maxresults=10"
	url := NewAutosuggestionService().SetHost("http://aa.bb").SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestAutosuggestionUrlWithoutHost(t *testing.T) {
	expected := "https://autocomplete.geocoder.api.here.com/6.2/suggest.json?app_id=xx&app_code=yy&query=sdf+dfdf+qwe&maxresults=10"
	url := NewAutosuggestionService().SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestAutosuggestionUrlWithMaxResult(t *testing.T) {
	expected := "http://aa.bb/6.2/suggest.json?app_id=xx&app_code=yy&query=sdf+dfdf+qwe&maxresults=100"
	url := NewAutosuggestionService().SetHost("http://aa.bb").SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").SetMaxResults(100).Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestAutosuggestionUrlWithOnlyLatitude(t *testing.T) {
	expected := "http://aa.bb/6.2/suggest.json?app_id=xx&app_code=yy&query=sdf+dfdf+qwe&maxresults=10"
	url := NewAutosuggestionService().SetHost("http://aa.bb").SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").SetLatitude(22.2222).Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestAutosuggestionUrlWithOnlyLongitude(t *testing.T) {
	expected := "http://aa.bb/6.2/suggest.json?app_id=xx&app_code=yy&query=sdf+dfdf+qwe&maxresults=10"
	url := NewAutosuggestionService().SetHost("http://aa.bb").SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").SetLongitude(33.333).Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestAutosuggestionUrlWithCoordinates(t *testing.T) {
	expected := "http://aa.bb/6.2/suggest.json?app_id=xx&app_code=yy&query=sdf+dfdf+qwe&maxresults=10&prox=11.111,22.222"
	url := NewAutosuggestionService().SetHost("http://aa.bb").SetAppID("xx").SetAppToken("yy").SetSearchPhrase("sdf dfdf qwe").SetLatitude(11.111).SetLongitude(22.222).Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}
