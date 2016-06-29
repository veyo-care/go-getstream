package getstream

type Feed struct {
	Client Client
	Slug   Slug
}

//func (f *Feed) Slug() Slug { return f.Slug }

func (f *Feed) Secret() string { return f.Client.Secret() }

func (f *Feed) AddActivity(activity *Activity) (*Activity, error) {
	activity = SignActivity(f.Secret(), activity)

	result := &Activity{}
	e := f.Client.Post(result, f.url(), f.Slug, activity)
	return result, e
}

func (f *Feed) AddActivities(activities []*Activity) error {
	signeds := make([]*Activity, len(activities), len(activities))
	for i, activity := range activities {
		signeds[i] = SignActivity(f.Secret(), activity)
	}

	// TODO: A result type to recieve the listing result.

	panic("not yet implemented.")
}

func (f *Feed) Activities(opt *Options) ([]*Activity, error) {
	result := ActivitiesResult{}
	e := f.Client.Get(&result, f.url(), f.Slug, opt)
	return result.Results, e
}

func (f *Feed) RemoveActivity(id string) error {
	return f.Client.Del(f.url()+id+"/", f.Slug)
}

func (f *Feed) Follow(feed, id string) error {

	followedSlug := Slug{feed, id, ""}
	signedFollowingSlug := SignSlug(f.Secret(), f.Slug)
	data := Follow{Target: followedSlug.String()}
	e := f.Client.Post(nil, f.url()+"following/", signedFollowingSlug, data)
	return e
}

func (f *Feed) Unfollow(feed, id string) error {
	return f.Client.Del(f.url()+"following/"+feed+":"+id+"/", f.Slug)
}

func (f *Feed) Following(opt *Options) ([]*Feed, error) {
	result := FollowersResult{}
	e := f.Client.Get(&result, f.url()+"following/", f.Slug, nil)
	return result.Results, e
}

func (f *Feed) Followers(opt *Options) ([]*Feed, error) {
	result := FollowersResult{}
	e := f.Client.Get(&result, f.url()+"followers/", f.Slug, nil)
	return result.Results, e
}

func (f *Feed) url() string {
	return "feed/" + f.Slug.Slug + "/" + f.Slug.ID + "/"
}

func (f *Feed) Empty() error {
	activities, e := f.Activities(nil)

	if e == nil {
		for _, activity := range activities {
			f.RemoveActivity(activity.ID)
		}
	}

	return e
}
