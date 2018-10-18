package tomtom

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const baseUrl = "https://api.tomtom.com"

type Client struct {
	BaseUrl     *url.URL
	ApiVersion  uint32
	ApiKey      string
	ContentType string
	httpClient  *http.Client

	Routing *RoutingService
}

func NewClient(apiVersion uint32, apiKey string, contentType string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, _ := url.Parse(baseUrl)
	c := &Client{
		BaseUrl:     u,
		ApiVersion:  apiVersion,
		ApiKey:      apiKey,
		ContentType: contentType,
		httpClient:  httpClient,
	}

	c.Routing = &RoutingService{client: c}
	return c
}

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	rel := &url.URL{Path: path}
	rel = c.BaseUrl.ResolveReference(rel)
	q := rel.Query()
	q.Set("key", c.ApiKey)
	rel.RawQuery = q.Encode()

	req, err := http.NewRequest(method, rel.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
