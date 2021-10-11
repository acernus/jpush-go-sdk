package jpush

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
}
func teardown() {

}

func TestPush(t *testing.T) {
	t.Skip()
	appKey, masterSecret := "", ""
	client := NewClient(appKey, masterSecret)
	msgID, err := client.Push(&PushReq{
		Platform: NewPlatform().UsePlatform(PlatformiOS),
		Audience: NewAudience().RegistrationList("", ""),
		Notification: NewNotification().IOS(&NotificationIOS{
			Alert: &IosPayload{
				Title: "iOS Test Title😁",
				Body:  "iOS test description",
			},
		}).Android(&NotificationAndroid{
			Title: "Android Test Title😁",
			Alert: "Android test description",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("msg id: %s", msgID)
}

func TestGetCid(t *testing.T) {
	t.Skip()
	appKey, masterSecret := "", ""
	client := NewClient(appKey, masterSecret)
	cidList, err := client.GetCID(&CidReq{
		Count: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	for _, item := range cidList {
		t.Log(item)
	}
}
