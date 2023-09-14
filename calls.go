package textgrid

import "fmt"

type Call struct {
	Sid            string       `json:"sid"`
	DateCreated    string       `json:"date_created"`
	DateUpdated    string       `json:"date_updated"`
	AccountSid     string       `json:"account_sid"`
	Status         string       `json:"status"`
	Duration       int          `json:"duration"`
	APIVersion     string       `json:"api_version"`
	Price          string       `json:"price"`
	PriceUnit      string       `json:"price_unit"`
	URI            string       `json:"uri"`
	From           string       `json:"from"`
	To             string       `json:"to"`
	QueueTime      TextGridTime `json:"queue_time"`
	Direction      string       `json:"direction"`
	StartTime      TextGridTime `json:"start_time"`
	EndTime        TextGridTime `json:"end_time"`
	PhoneNumberSid string       `json:"phone_number_sid"`
}

// GetCall gets the call details by the provided id
func (t *textGrid) GetCall(id string) (*Call, error) {
	result := new(Call)

	//Accounts/kjuUgB7bAst7NP5425662JHOC09Q==/Calls/CAIGVOmFk1Rj~vlCEDKnVBNuQ==.json
	endpoint := fmt.Sprintf("Accounts/%s/Calls/%s.json", t.AccountSid, id)

	if err := t.get(endpoint, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
