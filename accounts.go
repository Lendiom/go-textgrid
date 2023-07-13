package textgrid

import (
	"net/url"
)

type Account struct {
	Sid             string            `json:"sid"`
	FriendlyName    string            `json:"friendly_name"`
	Type            string            `json:"type"`
	AuthToken       string            `json:"auth_token"`
	WebhookSecret   string            `json:"webhook_secret"`
	OwnerAccountSid string            `json:"owner_account_sid"`
	DateCreated     TextGridTime      `json:"date_created"`
	DateUpdated     TextGridTime      `json:"date_updated"`
	Status          string            `json:"status"`
	SubresourceURIs map[string]string `json:"subresource_uris"`
	URI             string            `json:"uri"`
}

// CreateAccount creates a sub-account
func (t *textGrid) CreateAccount(data url.Values) (*Account, error) {
	result := new(Account)

	if err := t.postForm("Accounts.json", data, result); err != nil {
		return nil, err
	}

	return result, nil
}
