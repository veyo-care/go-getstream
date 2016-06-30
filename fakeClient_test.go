package getstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDel(t *testing.T) {
	activity := Activity{ID: "1", ForeignID: "10"}
	client := FakeConnect()
	feed := client.Feed("Patient", "2")
	_, e := feed.AddActivity(&activity)
	assert.Nil(t, e)
	activities, e := feed.Activities(nil)
	assert.Nil(t, e)
	assert.Len(t, activities, 1)
	e = feed.RemoveActivity("10")
	assert.Nil(t, e)
	activities, e = feed.Activities(nil)
	assert.Nil(t, e)
	assert.Len(t, activities, 0)
}
