package textgrid

import "fmt"

/*
{
  "action": "optedout",
  "date": "Tue, 28 May 2024 18:30:29 +0000",
  "verb": "STOP"
}
*/

// OptInStatusAction is the action that was taken. Either optedin, optedout, or nothing
type OptInStatusAction string

var (
	OptInStatusActionOptedOut OptInStatusAction = "optedout"
	OptInStatusActionOptedIn  OptInStatusAction = "optedin"
	OptInStatusActionNothing  OptInStatusAction = "nothing"
)

type OptInStatus struct {
	Action OptInStatusAction `json:"action"`
	Date   string            `json:"date,omitempty"`
	Verb   string            `json:"verb,omitempty"`
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
