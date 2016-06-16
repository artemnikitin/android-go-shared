package prototype

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/artemnikitin/android-go-shared/builder"
)

// HTTPDoer interface
type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

// Config represents config for HERE API client
type Config struct {
	appID    string
	appToken string
	client   HTTPDoer
}

// HereAPI
type HereAPI struct {
	appID    string
	appToken string
	client   *http.Client
}

// NewClient creates a new client for HERE API
func NewClient(config *Config) *HereAPI {
	api := &HereAPI{
		appID:    config.appID,
		appToken: config.appToken,
	}
	httpClient, ok := config.client.(*http.Client)
	if !ok {
		api.client = http.DefaultClient
	} else {
		api.client = httpClient
	}
	return api
}

func (a *HereAPI) GetPicture(params map[string]string) []byte {
	url := builder.NewMapTileService().
		SetAppID(a.appID).SetAppToken(a.appToken).
		SetParameters(params).NewBuild()
	resp, err := a.client.Get(url)
	if err != nil {
		return make([]byte, 0)
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		return make([]byte, 0)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return bytes
}
