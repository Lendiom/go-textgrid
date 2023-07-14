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
	FriendlyName string                      `json:"friendly_name"`
	PhoneNumber  string                      `json:"phone_number"`
	Lata         string                      `json:"lata"`
	RateCenter   string                      `json:"rate_center"`
	Region       string                      `json:"region"`
	IsoCountry   string                      `json:"iso_country"`
	Beta         bool                        `json:"beta"`
	Capabilities AvailableNumberCapabilities `json:"capabilities"`
}

type AvailableNumberCapabilities struct {
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
