package gitea

import (
	"encoding/json"
	"fmt"
)

// List the autenticated(or specified) user's followers
func (s *UsersService) ListFollowers(user string) ([]*User, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/followers", user)
	} else {
		u = "user/followers"
	}
	var users []*User
	req, err := s.client.getReq(u)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(req, &users)

	return users, nil
}

// List the users that the autenticated(or specified) user is following
func (s *UsersService) ListFollowing(user string) ([]*User, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/following", user)
	} else {
		u = "user/following"
	}
	var users []*User
	req, err := s.client.getReq(u)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(req, &users)

	return users, nil
}
