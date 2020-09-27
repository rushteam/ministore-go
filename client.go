package ministore

import (
	"context"
	"net/http"
	"net/url"
)

const (
	apiHost             = "https://api.weixin.qq.com"
	apiServiceCheckAuth = "/product/service/check_auth?component_access_token=%v"
	apiServiceList      = "/product/service/get_list?access_token=%v"
)

var Version = "0.0.0"

// Client type
type Client struct {
	host       *url.URL     // default host
	httpClient *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new client instance.
func New(secret, token string, options ...ClientOption) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if c.host == nil {
		u, err := url.ParseRequestURI(apiHost)
		if err != nil {
			return nil, err
		}
		c.host = u
	}
	return c, nil
}

// WithHTTPClient function
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = client
		return nil
	}
}

// WithHost function
func WithHost(host string) ClientOption {
	return func(c *Client) error {
		u, err := url.ParseRequestURI(host)
		if err != nil {
			return err
		}
		c.host = u
		return nil
	}
}
func (client *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+client.channelToken)
	req.Header.Set("User-Agent", "rushteam/ministore-Go/"+Version)
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return client.httpClient.Do(req)
}
