package github

import (
	"net/http"
	"net/url"
)

const DefaultBaseURL = "https://api.github.com"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL    string
	httpClient HTTPClient
}

func (c *Client) Request(method string, path string) (*http.Response, error) {
	baseURL := c.BaseURL
	httpClient := c.httpClient

	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	parsedBaseUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	url = parsedBaseUrl.ResolveReference(url)

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
