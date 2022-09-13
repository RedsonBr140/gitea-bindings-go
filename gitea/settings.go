package gitea

import "encoding/json"

type SettingsService service

// GeneralAPISettings contains global api settings exposed by it
type GeneralAPISettings struct {
	DefaultGitTreesPerPage int64 `json:"default_git_trees_per_page"`
	DefaultMaxBlobSize     int64 `json:"default_max_blob_size"`
	DefaultPagingNum       int64 `json:"default_paging_num"`
	MaxResponseItems       int64 `json:"max_response_items"`
}

// GeneralAttachmentSettings contains global Attachment settings exposed by API
type GeneralAttachmentSettings struct {
	AllowedTypes string `json:"allowed_types"`
	Enabled      bool   `json:"enabled"`
	MaxFiles     int64  `json:"max_files"`
	MaxSize      int64  `json:"max_size"`
}

// GeneralRepoSettings contains global repository settings exposed by API
type GeneralRepoSettings struct {
	HttpGitDisabled      bool `json:"http_git_disabled"`
	LfsDisabled          bool `json:"lfs_disabled"`
	MigrationsDisabled   bool `json:"migrations_disabled"`
	StarsDisabled        bool `json:"stars_disabled"`
	TimeTrackingDisabled bool `json:"time_tracking_disabled"`
}

// GeneralUISettings contains global ui settings exposed by API
type GeneralUISettings struct {
	AllowedReactions []string `json:"allowed_reactions"`
	CustomEmojis     []string `json:"custom_emojis"`
	DefaultTheme     string   `json:"default_theme"`
}

// Get instance's global settings for API.
func (s *SettingsService) GetAPISettings() (*GeneralAPISettings, error) {
	var APISettingsObj GeneralAPISettings

	resp, err := s.client.newRequest("GET", "settings/api", nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &APISettingsObj)
	return &APISettingsObj, nil
}

// Get intance's global settings for Attachment
func (s *SettingsService) GetAttachmentSettings() (*GeneralAttachmentSettings, error) {
	var AttachmentObj GeneralAttachmentSettings

	resp, err := s.client.newRequest("GET", "settings/attachment", nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &AttachmentObj)
	return &AttachmentObj, nil
}

// Get intance's global settings for repositorioes
func (s *SettingsService) GetRepositoriesSettings() (*GeneralRepoSettings, error) {
	var RepositoriesObj GeneralRepoSettings

	resp, err := s.client.newRequest("GET", "settings/repositories", nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &RepositoriesObj)
	return &RepositoriesObj, nil
}

// Get intance's global settings for UI.
func (s *SettingsService) GetUISettings() (*GeneralUISettings, error) {
	var UIObj GeneralUISettings

	resp, err := s.client.newRequest("GET", "settings/ui", nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &UIObj)
	return &UIObj, nil
}
