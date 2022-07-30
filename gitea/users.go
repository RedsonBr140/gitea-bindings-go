package gitea

import (
	"encoding/json"
	"fmt"
)

type UsersService service

type User struct {
	ID             int64  `json:"id,omitempty"`
	Login          string `json:"login,omitempty"`
	Full_Name      string `json:"full_name,omitempty"`
	Email          string `json:"email,omitempty"`
	AvatarURL      string `json:"avatar_url,omitempty"`
	Language       string `json:"language,omitempty"`
	IsAdmin        bool   `json:"is_admin,omitempty"`
	LastLogin      string `json:"last_login,omitempty"`
	Created        string `json:"created,omitempty"`
	IsRestricted   bool   `json:"restricted,omitempty"`
	IsActive       bool   `json:"active,omitempty"`
	ProhibitLogin  bool   `json:"prohibit_login,omitempty"`
	Location       string `json:"location,omitempty"`
	Website        string `json:"website,omitempty"`
	Description    string `json:"description,omitempty"`
	Visibility     string `json:"visibility,omitempty"`
	FollowersCount int64  `json:"followers_count,omitempty"`
	FollowingCount int64  `json:"following_count,omitempty"`
	StarredRepos   int64  `json:"starred_repos_count,omitempty"`
	Username       string `json:"username,omitempty"`
	// TODO: Finish https://codeberg.org/api/v1/users/redson
}

// Get information about the authenticated (or specified) user
func (s *UsersService) Get(user string) (*User, error) {
	var userObj User
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v", user)
	} else {
		u = "user"
	}

	resp, err := s.client.getReq(u)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &userObj)
	return &userObj, nil
}
