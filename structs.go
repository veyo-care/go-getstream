package getstream

import (
	"time"
)

type Activity struct {
	ID         string     `json:"id,omitempty"`
	ActorData  ActorData  `json:"actor_data,omitempty"`
	Actor      Slug       `json:"actor"`
	Verb       string     `json:"verb"`
	ObjectData ObjectData `json:"object_data"`
	Object     Slug       `json:"object"`
	Target     *Slug      `json:"target,omitempty"`
	TargetData TargetData `json:"target_data,omitempty"`
	RawTime    string     `json:"time,omitempty"`
	To         []Slug     `json:"to,omitempty"`
	ForeignID  string     `json:"foreign_id,omitempty"`
}

type Follow struct {
	Target             string `json:"target"`
	ActivityCopyResult int    `json:"activity_copy_limit,omitempty"`
}

type ObjectData struct {
	ID         string `json:"id"`
	Name       string `json:"name,omitempty"`
	Text       string `json:"text,omitempty"`
	PictureURL string `json:"picture_url,omitempty"`
}

type TargetData struct {
	ID       string   `json:"id"`
	Name     string   `json:"name,omitempty"`
	Location Location `json:"location,omitempty"`
}

type Location struct {
	Name      string `json:"name,omitempty"`
	Longitude string `json:"long,omitempty"`
	Latitude  string `json:"lat,omitempty"`
}

type ActorData struct {
	Name     string `json:"name,omitempty"`
	NickName string `json:"lastname,omitempty"`
	ID       string `json:"id"`
}

type ActivitiesResult struct {
	Next        string      `json:"next,omitempty"`
	RawDuration string      `json:"duration,omitempty"`
	Results     []*Activity `json:"results,omitempty"`
}

type FollowersResult struct {
	Next        string  `json:"next,omitempty"`
	RawDuration string  `json:"duration,omitempty"`
	Results     []*Feed `json:"results,omitempty"`
}

type Options struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IdGTE string `json:"id_gte,omitempty"`
	IdGT  string `json:"id_gt,omitempty"`
	IdLTE string `json:"id_lte,omitempty"`
	IdLT  string `json:"id_lt,omitempty"`

	Feeds    []*Feed `json:"feeds,omitempty"`
	MarkRead bool    `json:"mark_read,omitempty"`
	MarkSeen bool    `json:"mark_seen,omitempty"`
}

type Notification struct {
	Data    *Update `json"data"`
	Channel string  `json:"channel"`
}

type Update struct {
	Deletes []*Activity
	Inserts []*Activity

	UnreadCount int
	UnseenCount int
	PublishedAt time.Time
}
