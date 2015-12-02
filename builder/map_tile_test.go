package builder

import (
	"log"
	"testing"
)

func TestMapTileUrl(t *testing.T) {
	expected := "http://example.com/mia/1.6/mapview?app_id=xxx&app_code=yyy&c=11.11,22.22&h=12&w=23&ppi=1&z=18&u=10"
	builder := NewMapTileService()
	builder = builder.SetHost("http://example.com").SetAppId("xxx").SetAppToken("yyy")
	url := builder.SetLatitude(11.11).SetLongitude(22.22).SetWidth(23).SetHeight(12).SetDpi(1).Build()
	if url != expected {
		log.Fatal(url)
		t.Error("Strings should be equal")
	}
}
