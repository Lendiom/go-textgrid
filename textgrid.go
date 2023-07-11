package textgrid

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("textgrid")

// LogStackTrace logs a stack trace for the given error.
func logStackTrace(err error) {
	buf := make([]byte, 0, 16384)
	n := runtime.Stack(buf, false)
	if err != nil {
		log.Errorf("Non-nil error %s; stack trace %s", err.Error(), buf[:n])
	} else {
		log.Errorf("Nil error; stack trace %s", buf[:n])
	}
}

type TextGrid interface {
	CreateBrand(brand CreateBrandPayload) (*Brand, error)
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

	if resp.StatusCode != 200 {
		err = fmt.Errorf("non-200 status code %d returned from %s with body %s", resp.StatusCode, fullURL, data)
		logStackTrace(err)
		json.Unmarshal(data, returnValue) // try, anyway -- in case the caller wants error info
		return err
	}

	return json.Unmarshal(data, returnValue)
}
