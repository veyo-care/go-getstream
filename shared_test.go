package getstream

import "os"

var (
	TestAPIKey    = os.Getenv("GETSTREAM_KEY")
	TestAPISecret = os.Getenv("GETSTREAM_SECRET")
	TestAppID     = os.Getenv("GETSTREAM_APPID")
)

func ConnectTestClient(region string) *Client {
	return Connect(TestAPIKey, TestAPISecret, TestAppID, region)
}

func NewTestActivity() *Activity {
	return &Activity{
		Actor:      Slug{"user", "john", ""},
		Verb:       "add",
		ObjectData: ObjectData{ID: "id", Name: "object"},
		Object:     Slug{"object", "id", ""},
	}
}
