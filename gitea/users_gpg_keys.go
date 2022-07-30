package gitea

import (
	"encoding/json"
	"fmt"
)

type GPGKey struct {
	ID                int64      `json:"id,omitempty"`
	PrimaryKeyID      string     `json:"primary_key_id,omitempty"`
	KeyID             string     `json:"key_id,omitempty"`
	PublicKey         string     `json:"public_key,omitempty"`
	Emails            []GPGEmail `json:"emails,omitempty"`
	Subkeys           []GPGKey   `json:"subkeys,omitempty"`
	CanSign           bool       `json:"can_sign,omitempty"`
	CanEncryptComms   bool       `json:"can_encrypt_comms,omitempty"`
	CanEncryptStorage bool       `json:"can_encrypt_storage,omitempty"`
	CanCertify        bool       `json:"can_certify,omitempty"`
	Verified          bool       `json:"verified,omitempty"`
	CreatedAt         string     `json:"created_at,omitempty"`
	ExpiresAt         string     `json:"expires_at,omitempty"`
}

type GPGEmail struct {
	Email    string `json:"email,omitempty"`
	Verified string `json:"verified,omitempty"`
}

// ListGPGKeys list the public GPG keys for a user. If
// you pass a empty string, will fetch keys for the
// authenticated user.
func (s *UsersService) ListGPGKeys(user string) ([]*GPGKey, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/gpg_keys", user)
	} else {
		u = "user/gpg_keys"
	}

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var keys []*GPGKey
	json.Unmarshal(req, &keys)

	return keys, nil
}

// GetGPGKey gets details for a single GPG key.
// Requires authentication.
func (s *UsersService) GetGPGKey(id int64) (*GPGKey, error) {
	u := fmt.Sprintf("user/gpg_keys/%v", id)

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	key := &GPGKey{}
	json.Unmarshal(req, &key)
	return key, nil
}
