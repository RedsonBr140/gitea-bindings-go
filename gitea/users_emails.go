package gitea

import "encoding/json"

type UserEmail struct {
	Email      string `json:"email,omitempty"`
	IsVerified bool   `json:"verified,omitempty"`
	IsPrimary  bool   `json:"primary,omitempty"`
}

// List the authenticated user's email addresses
func (s *UsersService) ListEmails() ([]*UserEmail, error) {
	u := "user/emails"
	var mailsArray []*UserEmail

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(req, &mailsArray)

	return mailsArray, nil
}
