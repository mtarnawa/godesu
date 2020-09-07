package godesu

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	SCHEME = "https"
	DOMAIN = "a.4cdn.org"
	IMG_DOMAIN = "i.4cdn.org"
)

type Client struct {
	raw *http.Client
	globalHeaders map[string]string
}

func NewClient() *Client {
	return &Client{
		raw: &http.Client{
			Timeout: time.Second * 60,
		},
		globalHeaders: map[string]string{
			"Accept": "application/json",
		},
	}
}

func (c *Client) Get(endpoint string, model interface{}) error {
	resp, err := c.raw.Get(URL(endpoint))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, model)
	}
	return ErrNotOK{resp.StatusCode}
}

func URL(endpoint string) string {
	u := url.URL{
		Scheme:     SCHEME,
		Host:       DOMAIN,
		Path:       endpoint,
	}
	return u.String()
}

func IMG(endpoint string) string {
	u := url.URL{
		Scheme: SCHEME,
		Host: IMG_DOMAIN,
		Path: endpoint,
	}
	return u.String()
}
