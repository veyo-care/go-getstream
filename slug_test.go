package getstream

import (
	"encoding/json"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

var _ json.Marshaler = Slug{}
var _ json.Unmarshaler = &Slug{}
var _ fmt.Stringer = Slug{}

func TestSlug(t *testing.T) {
	assert.Equal(t, "test:456", Slug{"test", "456", ""}.String())
	assert.Equal(t, "test:123 token", Slug{"test", "123", "token"}.String())
}

func TestSlug_Valid(t *testing.T) {
	valids := []Slug{
		Slug{"feed", "123", ""},
		Slug{"feed", "456", "token"},
	}

	for _, slug := range valids {
		assert.True(t, slug.Valid())
	}

	invalids := []Slug{
		Slug{"", "", ""},
		Slug{"feed", "", ""},
		Slug{"", "123", ""},
		Slug{"", "", "token"},
	}

	for _, slug := range invalids {
		assert.False(t, slug.Valid())
	}
}

func TestSlug_JSON(t *testing.T) {
	marshals := map[Slug]string{
		Slug{"slug", "123", ""}:          `"slug:123"`,
		Slug{"slug", "123", "signature"}: `"slug:123 signature"`,
	}

	for slug, str := range marshals {
		bytes, e := json.Marshal(slug)
		assert.NoError(t, e, "failed to marshal slug: "+slug.String())
		assert.Equal(t, str, string(bytes))
	}

	unmarshals := map[string]Slug{
		`"slug:123"`:                Slug{"slug", "123", ""},
		`"slug:123 signature"`:      Slug{"slug", "123", "signature"},
		`["slug:123", "signature"]`: Slug{"slug", "123", "signature"},
	}

	for str, slug := range unmarshals {
		result := Slug{}
		e := json.Unmarshal([]byte(str), &result)
		assert.NoError(t, e, "failed to unmarshal json: "+str)
		assert.Equal(t, slug, result)
	}
}
