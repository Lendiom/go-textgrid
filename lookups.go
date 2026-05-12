package textgrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

var (
	ErrNumberNotFound = errors.New("number not found")
)

type Lookup struct {
	CallerName     *CallerName `json:"caller_name,omitempty"`
	CountryCode    string      `json:"country_code,omitempty"`
	PhoneNumber    string      `json:"phone_number,omitempty"`
	NationalFormat string      `json:"national_format,omitempty"`
	Carrier        *Carrier    `json:"carrier,omitempty"`
	URL            string      `json:"url,omitempty"`
}

type CallerName struct {
	CallerName string     `json:"caller_name,omitempty"`
	CallerType CallerType `json:"caller_type,omitempty"`
}

type CallerType string

var (
	CallerTypeBlank    CallerType = ""
	CallerTypeConsumer CallerType = "CONSUMER"
	CallerTypeBusiness CallerType = "BUSINESS"
)

type Carrier struct {
	MobileCountryCode string      `json:"mobile_country_code,omitempty"`
	MobileNetworkCode string      `json:"mobile_network_code,omitempty"`
	Name              string      `json:"name,omitempty"`
	Type              CarrierType `json:"type,omitempty"`
}

type CarrierType string

var (
	CarrierTypeBlank    CarrierType = ""
	CarrierTypeLandline CarrierType = "landline"
	CarrierTypeMobile   CarrierType = "mobile"
	CarrierTypeVoip     CarrierType = "voip"
)

type NumberLookup interface {
	Get(number string) (Lookup, error)
}

type numberLookup struct {
	accountSid  string
	bearerToken string
}

func (t *textGrid) Lookups() NumberLookup {
	return &numberLookup{
		accountSid:  t.AccountSid,
		bearerToken: t.bearerToken,
	}
}

func (nl *numberLookup) Get(number string) (Lookup, error) {
	const op = "lookups.get"

	res := Lookup{}

	params := url.Values{}
	params.Add("type", "carrier")
	params.Add("type", "caller-name")

	fullURL := fmt.Sprintf("https://lookups.textgrid.com/v1/PhoneNumbers/%s", number) + queryParams(params)
	slog.Debug("textgrid GET",
		slog.String("component", component),
		slog.String("op", op),
		slog.String("url", fullURL),
	)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		slog.Error("textgrid request build failure",
			slog.String("component", component),
			slog.String("op", op),
			slog.String("url", fullURL),
			slog.String("method", "GET"),
			slog.Any("error", err),
		)
		return res, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", nl.bearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("textgrid request transport failure",
			slog.String("component", component),
			slog.String("op", op),
			slog.String("url", fullURL),
			slog.String("method", "GET"),
			slog.Any("error", err),
		)
		return res, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("textgrid response body read failure",
			slog.String("component", component),
			slog.String("op", op),
			slog.String("url", fullURL),
			slog.Any("error", err),
		)
		return res, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return res, ErrNumberNotFound
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("non-200 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
		slog.Debug("textgrid request returned non-2xx",
			slog.String("component", component),
			slog.String("op", op),
			slog.String("url", fullURL),
			slog.String("method", "GET"),
			slog.Int("status_code", resp.StatusCode),
			slog.String("body", string(data)),
		)
		json.Unmarshal(data, &res) // try, anyway -- in case the caller wants error info
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	return res, nil
}
