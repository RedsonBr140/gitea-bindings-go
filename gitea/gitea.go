package gitea

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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
	Settings      *SettingsService
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
	c.Settings = (*SettingsService)(&c.common)

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

	if method == "POST" || method == "PATCH" {
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

// Get requests doesn't have a Body, it's the URL itself.
// So i'm parsing the opts object and putting on the link.
func (c *Client) parseQuery(url string, obj any) string {

	var (
		queries []string
		fields  []string
	)

	if url[len(url)-1:] != "?" {
		url += "?"
	}

	FieldValue := reflect.ValueOf(obj)
	st := reflect.TypeOf(obj)
	numFields := st.NumField()

	for i := 0; i < numFields; i++ {
		querie := strings.Split(st.Field(i).Tag.Get("json"), ",")
		queries = append(queries, querie[0])
		fields = append(fields, st.Field(i).Name)
	}
	for i, v := range queries {
		if FieldValue.FieldByName(fields[i]).String() == "" {
			continue
		}

		url += fmt.Sprintf("%v=%v&", v, FieldValue.FieldByName(fields[i]))
	}
	return url
}

func contains(i []int, code int) bool {
	for _, v := range i {
		if v == code {
			return true
		}
	}
	return false
}
