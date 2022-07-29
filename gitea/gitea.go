package gitea

import (
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://codeberg.org/api/v1/"
	userAgent      = "giteapi-go"
)

type Client struct {
	base         *url.URL
	client       *http.Client
	userAgent    string
	common       service
	Repositories *RepoService
	Users        *UsersService
}

type ClientOptions struct {
	BaseURL   *url.URL
	UserAgent string
}

type service struct {
	client *Client
}

// NewClient returns a new Gitea API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, opts *ClientOptions) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{base: baseURL, client: http.DefaultClient, userAgent: userAgent}
	c.common.client = c

	c.Repositories = (*RepoService)(&c.common)
	c.Users = (*UsersService)(&c.common)

	if opts != nil {
		if opts.BaseURL != nil {
			c.base = opts.BaseURL
		}
		if opts.UserAgent != "" {
			c.userAgent = opts.UserAgent
		}
	}
	return c
}

func (c *Client) getReq(urlStr string) ([]byte, error) {
	resp, err := http.Get(c.base.String() + urlStr)
	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return responseData, nil
}
