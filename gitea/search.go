package gitea

import "encoding/json"

type SearchService service

type UserSearchResult struct {
	Data []User `json:"data"`
	Ok   bool   `json:"ok"`
}

type UserSearchOpts struct {
	// Keyword
	Keyword string `json:"q"`
	Uid     int64  `json:"uid,omitempty"`
	Page    int    `json:"page,omitempty"`
	Limit   int    `json:"limit,omitempty"`
}

type RepoSearchResult struct {
	Data []Repo `json:"data"`
	Ok   bool   `json:"ok"`
}

type RepoSearchOpts struct {
	// Keyword
	Keyword string `json:"q"`
	// Limit search to repositories with keyword as topic
	Topic bool `json:"topic"`
	// include search of keyword within repository description
	IncludeDesc bool `json:"includeDesc"`
	// search only for repos that the user with the given id owns or contributes to
	Uid int64 `json:"uid"`
	// repo owner to prioritize in the results
	PriorityOwnerID int64 `json:"priority_owner_id"`
	// search only for repos that belong to the given team id
	TeamID int64 `json:"team_id"`
	// search only for repos that the user with the given id has starred
	StarredBy int64 `json:"starredBy"`
	// include private repositories this user has access to (defaults to true)
	Private bool `json:"private"`
	// show only pubic, private or all repositories (defaults to all)
	IsPrivate bool `json:"is_private"`
	// Include template repositories this user has access to (defaults to true)
	Template bool `json:"template"`
	// show only archived, non-archived or all repositories (defaults to all)
	Archived bool `json:"archived"`
	// type of repository to search for. Supported values are "fork", "source", "mirror" and "collaborative"
	Mode string `json:"mode"`
	// if uid is given, search only for repos that the user owns
	Exclusive bool `json:"exclusive"`
	// sort repos by attribute. Supported values are "alpha", "created", "updated", "size", and "id". Default is "alpha"
	Sort string `json:"sort"`
	// sort order, either "asc" (ascending) or "desc" (descending). Default is "asc", ignored if "sort" is not specified.
	Order string `json:"order"`
	// page number of results to return (1-based)
	Page int `json:"page"`
	// page size of results
	Limit int `json:"limit"`
}

// Search for users.
// if error, returns nil and the error.
func (s *SearchService) Users(opts UserSearchOpts) (*UserSearchResult, error) {
	var UserSearchObj UserSearchResult
	url := s.client.parseQuery("users/search", opts)

	resp, err := s.client.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &UserSearchObj)
	return &UserSearchObj, nil
}

// Search for repositories.
// if error, returns nil and the error.
func (s *SearchService) Repositories(opts RepoSearchOpts) (*RepoSearchResult, error) {
	var RepoSearchObj RepoSearchResult
	url := s.client.parseQuery("repos/search", opts)

	resp, err := s.client.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &RepoSearchObj)
	return &RepoSearchObj, nil
}
