package textgrid

import "fmt"

/*
{
  "action": "optedout",
  "date": "Tue, 28 May 2024 18:30:29 +0000",
  "verb": "STOP"
}
*/

type OptInStatus struct {
	Action string `json:"action"`
	Date   string `json:"date"`
	Verb   string `json:"verb"`
}

// GetMessage gets the message details by the provided id
func (t *textGrid) GetOptInStatus(to, from string) (*OptInStatus, error) {
	result := new(OptInStatus)

	//optin/status/+19132401531/+16095363466
	endpoint := fmt.Sprintf("optin/status/%s/%s", to, from)

	if err := t.get(endpoint, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
