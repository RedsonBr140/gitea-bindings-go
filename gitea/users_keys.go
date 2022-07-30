package gitea

import (
	"encoding/json"
	"fmt"
)

type SSHKey struct {
	ID        int64  `json:"id,omitempty"`
	Key       string `json:"key,omitempty"`
	Url       string `json:"url,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	KeyType   string `json:"key_type,omitempty"`
	User      User   `json:"user,omitempty"`
}
// ListKeys list the public SSH keys for a user. If
// you pass a empty string, will fetch keys for the
// authenticated user.
func (s *UsersService) ListKeys(user string) ([]*SSHKey, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/keys", user)
	} else {
		u = "user/keys"
	}

	var keys []*SSHKey
	req, err := s.client.getReq(u)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(req, &keys)
	return keys, nil
}

// GetKey gets details for a single SSH key.
// Requires authentication.
func (s *UsersService) GetKey(id int64) (*SSHKey, error) {
  u := fmt.Sprintf("user/keys/%v", id)
  key := &SSHKey{}

  req, err := s.client.getReq(u)
  if err != nil {
    return nil, err
  }

  json.Unmarshal(req, &key)
  return key, nil
}

