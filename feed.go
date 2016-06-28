package getstream

type Feed struct {
	Client
	slug Slug
}

func (f *Feed) Slug() Slug { return f.slug }

func (f *Feed) Secret() string { return f.Client.Secret() }

func (f *Feed) AddActivity(activity *Activity) (*Activity, error) {
	activity = SignActivity(f.Secret(), activity)

	result := &Activity{}
	e := f.post(result, f.url(), f.slug, activity)
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
	e := f.get(&result, f.url(), f.slug, opt)
	return result.Results, e
}

func (f *Feed) RemoveActivity(id string) error {
	return f.del(f.url()+id+"/", f.slug)
}

func (f *Feed) Follow(feed, id string) error {

	followedSlug := Slug{feed, id, ""}
	signedFollowingSlug := SignSlug(f.Secret(), f.slug)
	data := Follow{Target: followedSlug.String()}
	e := f.post(nil, f.url()+"following/", signedFollowingSlug, data)
	return e
}

func (f *Feed) Unfollow(feed, id string) error {
	return f.del(f.url()+"following/"+feed+":"+id+"/", f.slug)
}

func (f *Feed) Following(opt *Options) ([]*Feed, error) {
	result := FollowersResult{}
	e := f.get(&result, f.url()+"following/", f.slug, nil)
	return result.Results, e
}

func (f *Feed) Followers(opt *Options) ([]*Feed, error) {
	result := FollowersResult{}
	e := f.get(&result, f.url()+"followers/", f.slug, nil)
	return result.Results, e
}

func (f *Feed) url() string {
	return "feed/" + f.slug.Slug + "/" + f.slug.ID + "/"
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
