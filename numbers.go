package textgrid

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

type AvailableNumbersSearch struct {
	PostalCode string `url:"inPostalCode,omitempty"`
	Region     string `url:"inRegion,omitempty"`
	AreaCode   string `url:"areaCode,omitempty"`
	Contains   string `url:"contains,omitempty"`
	Lata       string `url:"inLata,omitempty"`
	RateCenter string `url:"inRateCenter,omitempty"`
}

type AvailableNumbers struct {
	AvailablePhoneNumbers []AvailableNumber `json:"available_phone_numbers"`
	URI                   string            `json:"uri"`
}

type AvailableNumber struct {
	FriendlyName string             `json:"friendly_name"`
	PhoneNumber  string             `json:"phone_number"`
	Lata         string             `json:"lata"`
	RateCenter   string             `json:"rate_center"`
	Region       string             `json:"region"`
	IsoCountry   string             `json:"iso_country"`
	Beta         bool               `json:"beta"`
	Capabilities NumberCapabilities `json:"capabilities"`
}

type NumberCapabilities struct {
	Voice bool `json:"voice"`
	Sms   bool `json:"sms"`
	Mms   bool `json:"mms"`
}

func (t *textGrid) ListAvailablePhoneNumbers(countryCode CountryCode, search AvailableNumbersSearch) (*AvailableNumbers, error) {
	result := new(AvailableNumbers)

	vals, err := query.Values(search)
	if err != nil {
		return nil, err
	}

	urlPartial := fmt.Sprintf("Accounts/%s/AvailablePhoneNumbers/%s/Local.json", t.AccountSid, countryCode)

	if err := t.get(urlPartial, vals, result); err != nil {
		return nil, err
	}

	return result, nil
}

type AddIncomingPhoneNumberPayload struct {
	PhoneNumber          string            `url:"phoneNumber"`
	FriendlyName         string            `url:"friendlyName,omitempty"`
	VoiceURL             string            `url:"voiceUrl,omitempty"`
	VoiceMethod          WebhookHTTPMethod `url:"voiceMethod,omitempty"`
	VoiceFallbackURL     string            `url:"voiceFallbackUrl,omitempty"`
	VoiceFallbackMethod  WebhookHTTPMethod `url:"voiceFallbackMethod,omitempty"`
	SmsURL               string            `url:"smsUrl,omitempty"`
	SmsMethod            WebhookHTTPMethod `url:"smsMethod,omitempty"`
	SmsFallbackURL       string            `url:"smsFallbackUrl,omitempty"`
	SmsFallbackMethod    WebhookHTTPMethod `url:"smsFallbackMethod,omitempty"`
	StatusCallbackURL    string            `url:"statusCallback,omitempty"`
	StatusCallbackMethod WebhookHTTPMethod `url:"statusCallbackMethod,omitempty"`
}

type IncomingPhoneNumber struct {
	ID                   string             `json:"sid"`
	AccountSid           string             `json:"account_sid"`
	FriendlyName         string             `json:"friendly_name"`
	PhoneNumber          string             `json:"phone_number"`
	VoiceURL             string             `json:"voice_url"`
	VoiceMethod          string             `json:"voice_method"`
	VoiceFallbackURL     string             `json:"voice_fallback_url"`
	VoiceFallbackMethod  string             `json:"voice_fallback_method"`
	DateCreated          string             `json:"date_created"`
	DateUpdated          string             `json:"date_updated"`
	SmsURL               string             `json:"sms_url"`
	SmsMethod            string             `json:"sms_method"`
	SmsFallbackURL       string             `json:"sms_fallback_url"`
	SmsFallbackMethod    string             `json:"sms_fallback_method"`
	Beta                 bool               `json:"beta"`
	Capabilities         NumberCapabilities `json:"capabilities"`
	StatusCallback       string             `json:"status_callback"`
	StatusCallbackMethod string             `json:"status_callback_method"`
	APIVersion           string             `json:"api_version"`
	URI                  string             `json:"uri"`
	Status               string             `json:"status"`
}

func (t *textGrid) AddIncomingPhoneNumber(payload AddIncomingPhoneNumberPayload) (*IncomingPhoneNumber, error) {
	result := new(IncomingPhoneNumber)

	vals, err := query.Values(payload)
	if err != nil {
		return nil, err
	}

	urlPartial := fmt.Sprintf("Accounts/%s/IncomingPhoneNumbers.json", t.AccountSid)

	if err := t.postForm(urlPartial, vals, result); err != nil {
		return nil, err
	}

	return result, nil
}
