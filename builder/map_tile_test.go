package builder

import (
	"log"
	"testing"
)

func TestMapTileUrl(t *testing.T) {
	expected := "http://example.com/mia/1.6/mapview?app_id=xxx&app_code=yyy&c=11.11,22.22&z=18&u=10&w=23&h=12&ppi=1"
	url := NewMapTileService().SetHost("http://example.com").SetAppID("xxx").SetAppToken("yyy").
		SetLatitude(11.11).SetLongitude(22.22).
		SetWidth(23).SetHeight(12).SetDpi(1).Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestMapTileUrlWithoutHost(t *testing.T) {
	expected := "https://image.maps.api.here.com/mia/1.6/mapview?app_id=xxx&app_code=yyy&c=11.11,22.22&z=18&u=10&w=23&h=12&ppi=1"
	url := NewMapTileService().SetAppID("xxx").SetAppToken("yyy").
		SetLatitude(11.11).SetLongitude(22.22).
		SetWidth(23).SetHeight(12).SetDpi(1).Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}

func TestMapTileUrlMinParams(t *testing.T) {
	expected := "https://image.maps.api.here.com/mia/1.6/mapview?app_id=&app_code=&c=0,0&z=18&u=10"
	url := NewMapTileService().Build()
	if url != expected {
		log.Println(url)
		t.Error("Strings should be equal")
	}
}
