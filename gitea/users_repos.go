package gitea

import (
	"encoding/json"
	"errors"
	"fmt"
)

type trust_model string

const (
	// Use the default repository trust model for this installation.
	TrustDefault trust_model = "default"

	// Trust Signatures by collaborators.
	// Valid signatures by collaborators of this repository will be marked "trusted" - (whether they match the committer or not).
	// Otherwise, valid signatures will be marked "untrusted" if the signature matches the committer and "unmatched" if not.
	TrustCollaborator = "collaborator"

	// Trust signatures that match committers (This matches GitHub and will force Gitea signed commits to have Gitea as committer)
	// Valid signatures will only be marked "trusted" if they match the committer, otherwise they will be marked "unmatched".
	// This will force Gitea to be the committer on signed commits with the actual committer marked as Co-authored-by: and Co-committed-by: trailer in the commit.
	// The default Gitea key must match a User in the database.
	TrustCommitter = "committer"

	// Trust Signatures by collaborators which matches the committer
	// Valid signatures by collaborators of this repository will be marked "trusted" if they match the committer.
	//Otherwise, valid signatures will be marked "untrusted" if the signature matches the committer and "unmatched" otherwise.
	// This will force Gitea to be marked as the committer on signed commits with the actual committer marked as Co-Authored-By: and Co-Committed-By: trailer in the commit.
	// The default Gitea key must match a User in the database.
	TrustCollaboratorcommitter = "collaboratorcommitter"
)

type CreateRepositoryOpts struct {
	// Whether the repository should be auto-initialized?
	AutoInit bool `json:"auto_init,omitempty"`
	// DefaultBranch of the repository (used when initializes and in template)
	DefaultBranch string `json:"default_branch,omitempty"`
	// Description of the repository to create
	Description string `json:"description,omitempty"`
	// Gitignores to use
	Gitignores string `json:"gitignores,omitempty"`
	// Label-Set to use
	IssueLabels string `json:"issue_labels,omitempty"`
	// License to use
	License string `json:"license,omitempty"`
	// Good repository names use short, memorable and unique keywords.
	Name string `json:"name"`

	// Only the owner or the organization members if they have rights, will be able to see it.
	Private bool `json:"private,omitempty"`
	//Readme of the repository to create.
	Readme string `json:"readme,omitempty"`

	// A repository to use as template for this one.
	Template string `json:"template,omitempty"`
	// Signature trust model
	TrustModel trust_model `json:"trust_model,omitempty"`
}

// ListRepos list the repos that a user owns. If
// you pass a empty string, will fetch keys for the
// authenticated user.
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

// Create a new repository for the authenticated user.
func (s *RepoService) CreateRepository(opts *CreateRepositoryOpts) ([]byte, error) {
	if opts == nil {
		return nil, errors.New("opts can't be nil")
	}
	reqBody, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.newRequest("POST", "user/repos", reqBody)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
