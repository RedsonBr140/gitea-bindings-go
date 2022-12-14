package gitea

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrRepoInfoMissing error = errors.New("repo owner or repo name is missing.")
)

type Repo struct {
	Name              string `json:"name,omitempty"`
	FullName          string `json:"full_name,omitempty"`
	Description       string `json:"description,omitempty"`
	HtmlURL           string `json:"html_url,omitempty"`
	SSHURL            string `json:"ssh_url,omitempty"`
	CloneURL          string `json:"clone_url,omitempty"`
	OriginalURL       string `json:"original_url,omitempty"`
	Website           string `json:"website,omitempty"`
	DefaultBranch     string `json:"default_branch,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	DefaultMergeStyle string `json:"default_merge_style,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	MirrorInterval    string `json:"mirror_interval,omitempty"`
	MirrorUpdated     string `json:"mirror_updated"`

	IsEmpty           bool `json:"empty,omitempty"`
	IsPrivate         bool `json:"private,omitempty"`
	IsFork            bool `json:"fork,omitempty"`
	IsTemplate        bool `json:"template,omitempty"`
	IsParent          bool `json:"parent,omitempty"`
	IsMirror          bool `json:"mirror,omitempty"`
	IsArchived        bool `json:"archived,omitempty"`
	HasIssues         bool `json:"has_issues,omitempty"`
	HasWiki           bool `json:"has_wiki,omitempty"`
	HasPullRequests   bool `json:"has_pull_requests,omitempty"`
	HasProjects       bool `json:"has_projects,omitempty"`
	IgnoreWhiteCon    bool `json:"ignore_whitespace_conflicts,omitempty"` // Ignore whitespace conflicts
	AllowMergeCommits bool `json:"allow_merge_commits,omitempty"`
	AllowRebase       bool `json:"allow_rebase,omitempty"`
	AllowReblaseExpli bool `json:"allow_rebase_explicit,omitempty"`
	AllowSquashMerge  bool `json:"allow_squash_merge,omitempty"`

	Size            int64 `json:"size,omitempty"`
	ID              int64 `json:"id,omitempty"`
	StarsCount      int64 `json:"stars_count,omitempty"`
	ForksCount      int64 `json:"forks_count,omitempty"`
	WatchersCount   int64 `json:"watchers_count,omitempty"`
	OpenIssuesCount int64 `json:"open_issues_count,omitempty"`
	OpenPRCounter   int64 `json:"open_pr_counter,omitempty"`
	ReleaseCounter  int64 `json:"release_counter,omitempty"`

	Owner       User      `json:"owner,omitempty"`
	Permissions RepoPerms `json:"permissions,omitempty"`

	// TODO: Implement the internal tracker
}

type RepoPerms struct {
	Admin bool `json:"admin,omitempty"`
	Push  bool `json:"push,omitempty"`
	Pull  bool `json:"pull,omitempty"`
}

type RepoService service

// Get information about the repository
func (s *RepoService) Get(repoOwner string, repoName string) (*Repo, error) {
	var RepoObj Repo
	var r string
	if repoOwner == "" || repoName == "" {
		return nil, ErrRepoInfoMissing
	}
	r = fmt.Sprintf("repos/%v/%v", repoOwner, repoName)

	resp, err := s.client.newRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &RepoObj)
	return &RepoObj, nil
}

func (s *RepoService) DeleteRepository(repoOwner string, repoName string) error {
	if repoName == "" {
		return errors.New("You need to specify a repository to delete")
	}
	u := fmt.Sprintf("repos/%v/%v", repoOwner, repoName)

	_, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return nil
}
