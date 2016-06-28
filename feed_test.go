package getstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeed(t *testing.T) {
	client := ConnectTestClient("eu-west")
	feed := client.Feed("user", "john2")
	activity := NewTestActivity()
	addedActivity, e := feed.AddActivity(activity)
	assert.Nil(t, e)
	assert.NotEqual(t, activity, addedActivity, "AddActivity should not modify existing instance.")
	assert.NotNil(t, addedActivity)
	assert.NotEmpty(t, addedActivity.ID)

	activities, e := feed.Activities(nil)
	assert.NoError(t, e)
	assert.NotEmpty(t, activities)
	assert.Len(t, activities, 1) // otherwise we might be getting result from another test run.
	assert.Equal(t, addedActivity.ID, activities[0].ID)

	e = feed.RemoveActivity(addedActivity.ID)
	assert.NoError(t, e)

	activities, e = feed.Activities(nil)
	assert.NoError(t, e)
	assert.Empty(t, activities)

}

func TestFollow(t *testing.T) {
	client := ConnectTestClient("eu-west")
	john := client.Feed("user", "john3")
	timeline := client.Feed("Test", "test3")
	activity := NewTestActivity()
	addedActivity, e := timeline.AddActivity(activity)
	assert.Nil(t, e)

	e = john.Follow("Test", "test3")
	assert.NoError(t, e)
	activities, e := john.Activities(nil)
	assert.NoError(t, e)
	assert.NotEmpty(t, activities)
	assert.Len(t, activities, 1) // otherwise we might be getting result from another test run.
	assert.Equal(t, addedActivity.ID, activities[0].ID)

	following, e := john.Following(nil)
	assert.NoError(t, e)
	assert.NotEmpty(t, following)
	assert.Len(t, following, 1) // otherwise we might be getting result from another test run.

	followers, e := timeline.Followers(nil)
	assert.NoError(t, e)
	assert.NotEmpty(t, followers)
	assert.Len(t, followers, 1) // otherwise we might be getting result from another test run.

	e = john.Unfollow("Test", "test3")
	assert.NoError(t, e)
	activities, e = john.Activities(nil)
	assert.NoError(t, e)
	assert.Empty(t, activities)

	following, e = john.Following(nil)
	assert.NoError(t, e)
	assert.Empty(t, following)

	followers, e = timeline.Followers(nil)
	assert.NoError(t, e)
	assert.Empty(t, followers)

	e = timeline.RemoveActivity(addedActivity.ID)
	assert.NoError(t, e)

}

func TestOptions(t *testing.T) {

}
