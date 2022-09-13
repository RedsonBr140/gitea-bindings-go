package gitea

import (
	"encoding/json"
	"errors"
)

type MiscellaneousService service

type MarkdownAsHTMLOpts struct {
	// Context to render
	Context string `json:"Context,omitempty"`
	// Mode to render
	Mode string `json:"Mode,omitempty"`
	// Text markdown to render. YOU HAVE TO SPECIFY.
	Text string `json:"Text"`
	// Is it a Wiki page?
	Wiki bool `json:"Wiki,omitempty"`
}
type Version struct {
	Version string `json:"version"`
}

// Returns the version of the gitea application.
func (s *MiscellaneousService) GetVersion() (*Version, error) {
	var version Version
	resp, err := s.client.newRequest("GET", "version", nil)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resp, &version)
	return &version, nil
}

// MarkdownAsHTML Render a markdown document as HTML.
// If error returns an empty string and the error.
func (s *MiscellaneousService) MarkdownAsHTML(opts *MarkdownAsHTMLOpts) (string, error) {
	if opts == nil {
		return "", errors.New("opts can't be nil.")
	}

	reqBody, err := json.Marshal(opts)
	if err != nil {
		return "", err
	}

	resp, err := s.client.newRequest("POST", "markdown", reqBody)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

// RawMarkdownAsHTML Render raw markdown as HTML.
// If error returns an empty string and the error.
func (s *MiscellaneousService) RawMarkdownAsHTML(markdown string) (string, error) {
	resp, err := s.client.newRequest("POST", "markdown/raw", []byte(markdown))
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
