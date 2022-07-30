package gitea

import (
	"encoding/json"
	"fmt"
)

func (s *UsersService) ListRepos(user string) ([]*Repo, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/repos", user)
	} else {
		u = "user/repos"
	}
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	repos := []*Repo{}

	json.Unmarshal(req, &repos)
	return repos, nil
}
