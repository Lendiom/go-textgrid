package textgrid

import "fmt"

type IncomingMessage struct {
	AccountSID        string `form:"AccountSid"`
	APIVersion        string `form:"ApiVersion"`
	Body              string `form:"Body"`
	From              string `form:"From"`
	MessageSID        string `form:"MessageSid"`
	NumberOfMedia     int    `form:"NumMedia"`
	NumberOfSegments  int    `form:"NumSegments"`
	MediaUrl0         string `form:"MediaUrl0"`
	MediaContentType0 string `form:"MediaContentType0"`
	MediaUrl1         string `form:"MediaUrl1"`
	MediaContentType1 string `form:"MediaContentType1"`
	MediaUrl2         string `form:"MediaUrl2"`
	MediaContentType2 string `form:"MediaContentType2"`
	MediaUrl3         string `form:"MediaUrl3"`
	MediaContentType3 string `form:"MediaContentType3"`
	SmsMessageSid     string `form:"SmsMessageSid"`
	SmsSID            string `form:"SmsSid"`
	SmsStatus         string `form:"SmsStatus"`
	To                string `form:"To"`
}

type Message struct {
	AccountSID      string                `json:"account_sid"`
	SID             string                `json:"sid"`
	Body            string                `json:"body"`
	From            string                `json:"from"`
	To              string                `json:"to"`
	Status          string                `json:"status"`
	Direction       string                `json:"direction"`
	Price           string                `json:"price"`
	Surcharge       string                `json:"surcharge"`
	PriceUnit       string                `json:"price_unit"`
	NumSegments     string                `json:"num_segments"`
	NumMedia        string                `json:"num_media"`
	Email           bool                  `json:"email"`
	DateCreated     TextGridTime          `json:"date_created"`
	DateSent        TextGridTime          `json:"date_sent"`
	DateUpdated     TextGridTime          `json:"date_updated"`
	CarrierNetwork  string                `json:"carrierNetwork"`
	MessageClass    string                `json:"messageClass"`
	APIVersion      string                `json:"api_version"`
	URI             string                `json:"uri"`
	SubresourceUris MessageSubresourceUri `json:"subresource_uris"`
}

type MessageSubresourceUri struct {
	Media string `json:"media"`
}

// GetMessage gets the message details by the provided id
func (t *textGrid) GetMessage(id string) (*Message, error) {
	result := new(Message)

	//Accounts/kjuUgB7bAst7NP5425662JHOC09Q==/Messages/CAIGVOmFk1Rj~vlCEDKnVBNuQ==.json
	endpoint := fmt.Sprintf("Accounts/%s/Messages/%s.json", t.AccountSid, id)

	if err := t.get(endpoint, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
