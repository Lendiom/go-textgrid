package textgrid

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// http11Client forces HTTP/1.1 for endpoints that reject HTTP/2.
var http11Client = &http.Client{
	Transport: &http.Transport{
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	},
}

type CreateCampaignPayload struct {
	BrandID            string    `json:"brandId"`
	Usecase            UseCase   `json:"usecase"`
	SubUsecases        []UseCase `json:"subUsecases"`
	Description        string    `json:"description"`
	EmbeddedLink       bool      `json:"embeddedLink"`
	EmbeddedPhone      bool      `json:"embeddedPhone"`
	NumberPool         bool      `json:"numberPool"`
	AgeGated           bool      `json:"ageGated"`
	DirectLending      bool      `json:"directLending"`
	SubscriberOptin    bool      `json:"subscriberOptin"`
	SubscriberOptout   bool      `json:"subscriberOptout"`
	SubscriberHelp     bool      `json:"subscriberHelp"`
	Sample1            string    `json:"sample1,omitempty"`
	Sample2            string    `json:"sample2,omitempty"`
	Sample3            string    `json:"sample3,omitempty"`
	Sample4            string    `json:"sample4,omitempty"`
	Sample5            string    `json:"sample5,omitempty"`
	MessageFlow        string    `json:"messageFlow,omitempty"`
	HelpKeywords       string    `json:"helpKeywords,omitempty"`
	HelpMessage        string    `json:"helpMessage,omitempty"`
	OptinKeywords      string    `json:"optinKeywords,omitempty"`
	OptinMessage       string    `json:"optinMessage,omitempty"`
	OptoutKeywords     string    `json:"optoutKeywords,omitempty"`
	OptoutMessage      string    `json:"optoutMessage,omitempty"`
	ReferenceID        string    `json:"referenceId,omitempty"`
	AutoRenewal        bool      `json:"autoRenewal"`
	AffiliateMarketing bool      `json:"affiliateMarketing"`
}

type Campaign struct {
	ID                 string                    `json:"campaignId"`
	AccountSid         string                    `json:"accountSid"`
	BrandID            string                    `json:"brandId"`
	CspID              string                    `json:"cspId"`
	ResellerID         string                    `json:"resellerId"`
	ReferenceID        string                    `json:"referenceId"`
	Status             string                    `json:"status"`
	DateCreated        TextGridTime              `json:"dateCreated"`
	BilledDate         TextGridTime              `json:"billedDate"`
	AutoRenewal        bool                      `json:"autoRenewal"`
	Usecase            UseCase                   `json:"usecase"`
	SubUsecases        []UseCase                 `json:"subUsecases"`
	Description        string                    `json:"description"`
	EmbeddedLink       bool                      `json:"embeddedLink"`
	EmbeddedPhone      bool                      `json:"embeddedPhone"`
	AffiliateMarketing bool                      `json:"affiliateMarketing"`
	NumberPool         bool                      `json:"numberPool"`
	AgeGated           bool                      `json:"ageGated"`
	DirectLending      bool                      `json:"directLending"`
	SubscriberOptin    bool                      `json:"subscriberOptin"`
	SubscriberOptout   bool                      `json:"subscriberOptout"`
	SubscriberHelp     bool                      `json:"subscriberHelp"`
	Sample1            string                    `json:"sample1,omitempty"`
	Sample2            string                    `json:"sample2,omitempty"`
	Sample3            string                    `json:"sample3,omitempty"`
	Sample4            string                    `json:"sample4,omitempty"`
	Sample5            string                    `json:"sample5,omitempty"`
	MessageFlow        string                    `json:"messageFlow"`
	HelpKeywords       string                    `json:"helpKeywords"`
	HelpMessage        string                    `json:"helpMessage"`
	OptinKeywords      string                    `json:"optinKeywords"`
	OptinMessage       string                    `json:"optinMessage"`
	OptoutKeywords     string                    `json:"optoutKeywords"`
	OptoutMessage      string                    `json:"optoutMessage"`
	DcaSharingStatus   SecondaryDcaSharingStatus `json:"SecondaryDcaSharingStatus"`
	DcaDeclineReason   string                    `json:"SecondaryDcaDeclineReason"`
}

type SecondaryDcaSharingStatus string

var (
	SecondaryDcaSharingStatusBlank    SecondaryDcaSharingStatus = ""
	SecondaryDcaSharingStatusPending  SecondaryDcaSharingStatus = "PENDING"
	SecondaryDcaSharingStatusAccepted SecondaryDcaSharingStatus = "ACCEPTED"
	SecondaryDcaSharingStatusDeclined SecondaryDcaSharingStatus = "DECLINED"
)

// CreateCampaign submits a campaign for registration and review
func (t *textGrid) CreateCampaign(payload CreateCampaignPayload) (*Campaign, error) {
	result := new(Campaign)

	if err := t.post("campaigns/campaign", payload, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetCampaigns returns all of the campaigns associated with the account
func (t *textGrid) GetCampaigns() ([]Campaign, error) {
	result := make([]Campaign, 0)

	if err := t.get("campaigns/campaign", nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetCampaign gets the campaign details by the provided id
func (t *textGrid) GetCampaign(id string) (*Campaign, error) {
	result := new(Campaign)

	if err := t.get("campaigns/campaign/"+id, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (t *textGrid) DeactivateCampaign(id string) error {
	return t.delete("campaigns/campaign/"+id, nil)
}

type attachNumberToCampaign struct {
	PhoneNumberSids []string `json:"phoneNumberSids"`
}

func (t *textGrid) AttachNumberToCampaign(id, numberID string) error {
	payload := attachNumberToCampaign{
		PhoneNumberSids: []string{numberID},
	}

	fullURL := t.generateFullUrl("campaigns/number/" + id)
	log.Debugf("TextGrid POST %s", fullURL)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", t.bearerToken)

	// Use HTTP/1.1 client — TextGrid's campaigns/number endpoint rejects HTTP/2.
	resp, err := http11Client.Do(req)
	if err != nil {
		logStackTrace(err)
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logStackTrace(err)
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("non-200/201 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
	}

	return nil
}
