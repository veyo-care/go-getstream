package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client interface {
	BaseURL() *url.URL
	Feed(slug, id string) *Feed
	get(result interface{}, path string, slug Slug, opt *Options) error
	post(result interface{}, path string, slug Slug, payload interface{}) error
	del(path string, slug Slug) error
	request(result interface{}, method, path string, slug Slug, payload interface{}) error
	absoluteUrl(path string) (result *url.URL, e error)
	Secret() string
}

type TrueClient struct {
	http    *http.Client
	baseURL *url.URL // https://api.getstream.io/api/

	key      string
	secret   string
	appID    string
	location string // https://location-api.getstream.io/api/
}

func (c TrueClient) Secret() string {
	return c.secret
}

func Connect(key, secret, appID, location string) TrueClient {
	baseURLStr := "https://api.getstream.io/api/v1.0/"
	if location != "" {
		baseURLStr = "https://" + location + "-api.getstream.io/api/v1.0/"
	}

	baseURL, e := url.Parse(baseURLStr)
	if e != nil {
		panic(e) // failfast, url shouldn't be invalid anyway.
	}

	return TrueClient{
		http:    &http.Client{},
		baseURL: baseURL,

		key:      key,
		secret:   secret,
		appID:    appID,
		location: location,
	}
}

func (c TrueClient) BaseURL() *url.URL { return c.baseURL }

func (c TrueClient) Feed(slug, id string) *Feed {
	return &Feed{
		Client: c,
		slug:   SignSlug(c.secret, Slug{slug, id, ""}),
	}
}

func (c TrueClient) get(result interface{}, path string, slug Slug, opt *Options) error {
	return c.request(result, "GET", path, slug, nil)
}

func (c TrueClient) post(result interface{}, path string, slug Slug, payload interface{}) error {
	return c.request(result, "POST", path, slug, payload)
}

func (c TrueClient) del(path string, slug Slug) error {
	return c.request(nil, "DELETE", path, slug, nil)
}

func (c TrueClient) request(result interface{}, method, path string, slug Slug, payload interface{}) error {
	absUrl, e := c.absoluteUrl(path)
	if e != nil {
		return e
	}

	buffer := []byte{}
	if payload != nil {
		if buffer, e = json.Marshal(payload); e != nil {
			return e
		}
	}

	req, e := http.NewRequest(method, absUrl.String(), bytes.NewBuffer(buffer))
	if e != nil {
		return e
	}

	req.Header.Set("Content-Type", "application/json")
	if slug.Token != "" {
		req.Header.Set("Authorization", slug.Signature())
	}

	resp, e := c.http.Do(req)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	buffer, e = ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}

	switch {
	case 200 <= resp.StatusCode && resp.StatusCode < 300: // SUCCESS
		if result != nil {
			if e = json.Unmarshal(buffer, result); e != nil {
				return e
			}
		}

	default:
		err := &Error{}
		if e = json.Unmarshal(buffer, err); e != nil {
			panic(e)
			return errors.New(string(buffer))
		}

		return err
	}

	return nil
}

func (c TrueClient) absoluteUrl(path string) (result *url.URL, e error) {
	if result, e = url.Parse(path); e != nil {
		return nil, e
	}

	// DEBUG: Use this line to send stuff to a proxy instead.
	// c.baseURL, _ = url.Parse("http://0.0.0.0:8000/")
	result = c.baseURL.ResolveReference(result)

	qs := result.Query()
	qs.Set("api_key", c.key)
	if c.location == "" {
		qs.Set("location", "unspecified")
	} else {
		qs.Set("location", c.location)
	}
	result.RawQuery = qs.Encode()

	return result, nil
}
