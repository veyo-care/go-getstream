package getstream

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

//here we use foreign_id as getstream_id

type FakeClient struct {
	values map[string]map[string]map[string]*Activity //type of feed, feed_id, activity_foreign_id
}

func (c FakeClient) Secret() string {
	return ""
}

func FakeConnect() *FakeClient {
	v := make(map[string]map[string]map[string]*Activity)
	v["Patient"] = make(map[string]map[string]*Activity) //only 2 types of feed
	v["Notification"] = make(map[string]map[string]*Activity)
	return &FakeClient{
		values: v,
	}
}

func (c FakeClient) BaseURL() *url.URL { return nil }

func newFeed(client FakeClient, slug Slug) *Feed {
	return &Feed{
		Client: client,
		Slug:   slug,
	}
}

func (c FakeClient) Feed(slug, id string) *Feed {
	return newFeed(c, Slug{slug, id, ""})
}

func (c FakeClient) Get(result interface{}, path string, slug Slug, opt *Options) error {
	parsedPath := parse(path)

	feeds, ok := c.values[parsedPath.feeds]
	if !ok {
		return fmt.Errorf("Could not find feeds for name %s", parsedPath.feeds)
	}
	activities, ok := feeds[parsedPath.feed]
	if !ok {
		return fmt.Errorf("Could not find activities for id %s", parsedPath.feed)
	}
	res := make([]*Activity, len(activities), len(activities))

	i := 0
	for _, value := range activities {
		res[i] = value
		i = i + 1
	}

	buffer, _ := json.Marshal(ActivitiesResult{Results: res})
	json.Unmarshal(buffer, result)

	return nil

}

func (c FakeClient) Post(result interface{}, path string, slug Slug, payload interface{}) error {
	parsedPath := parse(path)

	activity, ok := payload.(*Activity)
	activity.ID = activity.ForeignID
	if !ok {
		return fmt.Errorf("FakeClient can only receive activity")
	}
	feeds, ok := c.values[parsedPath.feeds]
	if !ok {
		return fmt.Errorf("Should post in Patient or Notification")
	}
	_, ok = feeds[parsedPath.feed]
	if !ok {
		newActivities := make(map[string]*Activity)
		newActivities[activity.ForeignID] = activity
		c.values[parsedPath.feeds][parsedPath.feed] = newActivities
	} else {
		c.values[parsedPath.feeds][parsedPath.feed][activity.ForeignID] = activity
	}
	return nil

}

func (c FakeClient) Del(path string, slug Slug) error {
	parsedPath := parse(path)
	delete(c.values[parsedPath.feeds][parsedPath.feed], parsedPath.ActivityID)
	return nil
}

type parsedPath struct {
	feeds      string
	feed       string
	ActivityID string
}

func parse(s string) parsedPath {
	split := strings.Split(s, "/")
	p := parsedPath{
		feeds: split[1],
		feed:  split[2],
	}
	if len(split) >= 4 {
		p.ActivityID = split[3]
	}
	return p
}
