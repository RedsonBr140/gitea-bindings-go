package gitea

import (
	"encoding/json"
	"errors"
)

type UserSettings struct {
	FullName      string `json:"full_name,omitempty"`
	Website       string `json:"website,omitempty"`
	Description   string `json:"description,omitempty"`
	Location      string `json:"location,omitempty"`
	Language      string `json:"language,omitempty"`
	Theme         string `json:"theme,omitempty"`
	DiffViewStyle string `json:"diff_view_style,omitempty"`
	HideEmail     bool   `json:"hide_email,omitempty"`
	HideActivity  bool   `json:"hide_activity,omitempty"`
}

// GetSettings Get the authenticated user settings.
// If a error occurs, it returns nil and the error.
func (s *UsersService) GetSettings() (*UserSettings, error) {
	var SettingsObj UserSettings
	req, err := s.client.newRequest("GET", "user/settings", nil)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(req, &SettingsObj)
	return &SettingsObj, nil
}

// EditSettings update the authenticated user settings.
// If a error occurs, it returns nil and the error.
func (s *UsersService) EditSettings(opts *UserSettings) (*UserSettings, error) {
	var usersettingsObj UserSettings
	if opts == nil {
		return nil, errors.New("opts can't be nil.")
	}
	reqBody, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.newRequest("PATCH", "user/settings", reqBody)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resp, &usersettingsObj)
	return &usersettingsObj, nil
}
