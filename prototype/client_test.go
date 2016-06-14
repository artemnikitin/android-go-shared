package prototype

import (
	"os"
	"testing"
)

var (
	appID    = os.Getenv("HERE_ANDROID_GO_ID")
	appToken = os.Getenv("HERE_ANDROID_GO_TOKEN")
)

func TestGetPicturePositive(t *testing.T) {
	client := NewClient(&Config{
		appID:    appID,
		appToken: appToken,
	})
	image := client.GetPicture(52.5308599, 13.38469, 320, 240, 24)
	if image[0] != 0xFF || image[1] != 0xD8 {
		t.Error("Should return an image in response")
	}
}
