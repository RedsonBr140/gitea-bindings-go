package gitea

import (
	"bytes"
	"encoding/json"
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
	base          *url.URL
	client        *http.Client
	userAgent     string
	token         string
	common        service
	Repositories  *RepoService
	Users         *UsersService
	Miscellaneous *MiscellaneousService
}

type ClientOptions struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string
}

type service struct {
	client *Client
}

type PostError struct {
	Message string `json:"message"`
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
	c.Miscellaneous = (*MiscellaneousService)(&c.common)

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

func (c *Client) newRequest(method string, urlStr string, reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest(method, c.base.String()+urlStr, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "token "+c.token)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		accepted := []int{200, 201}
		StatusIsOK := contains(accepted, resp.StatusCode)

		if !StatusIsOK {
			var postError PostError
			json.Unmarshal(responseData, &postError)
			return nil, errors.New(postError.Message)

		}
	}
	return responseData, nil
}

func contains(i []int, code int) bool {
	for _, v := range i {
		if v == code {
			return true
		}
	}
	return false
}
