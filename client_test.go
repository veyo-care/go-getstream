package getstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_BaseURL(t *testing.T) {
	locations := map[string]string{
		"":        "https://api.getstream.io/api/v1.0/",
		"us-east": "https://us-east-api.getstream.io/api/v1.0/",
		"xyz":     "https://xyz-api.getstream.io/api/v1.0/",
	}

	for location, url := range locations {
		client := ConnectTestClient(location)
		assert.Equal(t, url, client.BaseURL().String())
	}
}
