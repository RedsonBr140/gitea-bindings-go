package gitea

import (
	"bytes"
	"errors"
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
	token        string
	common       service
	Repositories *RepoService
	Users        *UsersService
}

type ClientOptions struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string
}

type service struct {
	client *Client
}

// NewClient returns a new Gitea API client.
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an ClientOptions.token that will perform the authentication.
func NewClient(opts *ClientOptions) *Client {
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
		if opts.Token != "" {
			c.token = opts.Token
		}
	}
	return c
}

func (c *Client) getReq(urlStr string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.base.String()+urlStr, nil)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "token "+c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return responseData, nil
}

func (c *Client) postReq(urlStr string, reqBody []byte) error {
	req, err := http.NewRequest("POST", c.base.String()+urlStr, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	if c.token == "" {
		return errors.New("The access token is missing...")
	}

	req.Header.Set("Authorization", "token "+c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
