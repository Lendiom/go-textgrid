package textgrid

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("textgrid")

// LogStackTrace logs a stack trace for the given error.
func logStackTrace(err error) {
	buf := make([]byte, 0, 16384)
	n := runtime.Stack(buf, false)
	if err == nil {
		log.Errorf("nil error; stack trace %s", buf[:n])
	} else {
		log.Errorf("non-nil error %s; stack trace %s", err.Error(), buf[:n])
	}
}

type TextGrid interface {
	CreateAccount(data url.Values) (*Account, error)

	CreateBrand(brand CreateBrandPayload) (*Brand, error)
	GetBrand(id string) (*Brand, error)
	DeleteBrand(id string) error

	CreateCampaign(payload CreateCampaignPayload) (*Campaign, error)
	GetCampaigns() ([]Campaign, error)
	GetCampaign(id string) (*Campaign, error)
	DeactivateCampaign(id string) error
	AttachNumberToCampaign(id, numberID string) error

	GetCall(id string) (*Call, error)
	InitiateCall(call CallInitiatePayload) (*Call, error)

	GetMessage(id string) (*Message, error)

	ListAvailablePhoneNumbers(countryCode CountryCode, search AvailableNumbersSearch) (*AvailableNumbers, error)
	AddIncomingPhoneNumber(payload AddIncomingPhoneNumberPayload) (*IncomingPhoneNumber, error)

	Lookups() NumberLookup
}

// Lob represents information on how to connect to the lob.com API.
type textGrid struct {
	BaseAPI     string
	AccountSid  string
	authToken   string
	bearerToken string
}

// Base URL and API version for Lob.
const (
	BaseAPI    = "https://api.textgrid.com"
	APIVersion = "2020-01-01"
)

// NewLob creates an object that can be used to connect to the lob.com API.
func NewTextGrid(baseAPI, accountSid, authToken string) *textGrid {
	token := fmt.Sprintf("%s:%s", accountSid, authToken)
	base64Value := base64.StdEncoding.EncodeToString([]byte(token))

	if baseAPI == "" {
		baseAPI = BaseAPI
	}

	return &textGrid{
		BaseAPI:     baseAPI,
		AccountSid:  accountSid,
		authToken:   authToken,
		bearerToken: fmt.Sprintf("Bearer %s", base64Value),
	}
}

func (t *textGrid) generateFullUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", t.BaseAPI, APIVersion, endpoint)
}

func (t *textGrid) post(endpoint string, payload, returnValue interface{}) error {
	fullURL := t.generateFullUrl(endpoint)
	log.Debugf("TextGrid POST %s", fullURL)

	body, err := json.Marshal(payload)
	if err != nil {
		log.Debugf("Failed to unmarshal the payload: %+v", payload)
		return err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(body))
	if err != nil {
		return err
	}

	if len(body) != 0 {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", t.bearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logStackTrace(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logStackTrace(err)
		return err
	}

	if resp.StatusCode == http.StatusBadRequest {
		var respErr genericError
		if err := json.Unmarshal(data, &respErr); err != nil {
			return fmt.Errorf("status code 400 return from %s with body: %s", fullURL, data)
		}

		var badRequestErr badRequestError
		if err := json.Unmarshal([]byte(respErr.Error), &badRequestErr); err != nil {
			return fmt.Errorf("status code 400 return from %s with body: %s", fullURL, data)
		}

		if badRequestErr.Field == "" {
			return fmt.Errorf("%s. For the url %s", badRequestErr.Description, fullURL)
		}

		return fmt.Errorf("the field %s is not correctly formatted. %s. For the url %s", badRequestErr.Field, badRequestErr.Description, fullURL)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("non-200/201 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
		logStackTrace(err)
		json.Unmarshal(data, returnValue) // try, anyway -- in case the caller wants error info
		return err
	}

	// not every endpoint returns something
	if returnValue == nil {
		return nil
	}

	return json.Unmarshal(data, returnValue)
}

func (t *textGrid) postForm(endpoint string, payload url.Values, returnValue interface{}) error {
	fullURL := t.generateFullUrl(endpoint)
	log.Debugf("TextGrid POST form %s", fullURL)

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", t.bearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logStackTrace(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logStackTrace(err)
		return err
	}

	if resp.StatusCode == http.StatusBadRequest {
		var respErr genericError
		if err := json.Unmarshal(data, &respErr); err != nil {
			return fmt.Errorf("status code 400 return from %s with body: %s", fullURL, data)
		}

		var badRequestErr badRequestError
		if err := json.Unmarshal([]byte(respErr.Error), &badRequestErr); err != nil {
			return fmt.Errorf("status code 400 return from %s with body: %s", fullURL, data)
		}

		if badRequestErr.Field == "" {
			return fmt.Errorf("%s. For the url %s", badRequestErr.Description, fullURL)
		}

		return fmt.Errorf("the field %s is not correctly formatted. %s. For the url %s", badRequestErr.Field, badRequestErr.Description, fullURL)
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("non-200 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
		logStackTrace(err)
		json.Unmarshal(data, returnValue) // try, anyway -- in case the caller wants error info
		return err
	}

	return json.Unmarshal(data, returnValue)
}

func (t *textGrid) get(endpoint string, params url.Values, returnValue interface{}) error {
	fullURL := t.generateFullUrl(endpoint) + queryParams(params)
	log.Debugf("TextGrid GET %s", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		logStackTrace(err)
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", t.bearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logStackTrace(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logStackTrace(err)
		return err
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("non-200 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
		logStackTrace(err)
		json.Unmarshal(data, returnValue) // try, anyway -- in case the caller wants error info
		return err
	}

	return json.Unmarshal(data, returnValue)
}

func (t *textGrid) delete(endpoint string, params url.Values) error {
	fullURL := t.generateFullUrl(endpoint) + queryParams(params)
	log.Debugf("TextGrid DELETE %s", fullURL)
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		logStackTrace(err)
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", t.bearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logStackTrace(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logStackTrace(err)
		return err
	}

	if resp.StatusCode != 204 {
		err = fmt.Errorf("non-204 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
		logStackTrace(err)
		return err
	}

	return nil
}
