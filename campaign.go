package textgrid

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
	ID                 string    `json:"campaignId"`
	BrandID            string    `json:"brandId"`
	CspID              string    `json:"cspId"`
	ResellerID         string    `json:"resellerId"`
	ReferenceID        string    `json:"referenceId"`
	Status             string    `json:"status"`
	CreateDate         string    `json:"createDate"`
	BilledDate         string    `json:"billedDate"`
	AutoRenewal        bool      `json:"autoRenewal"`
	Usecase            UseCase   `json:"usecase"`
	SubUsecases        []UseCase `json:"subUsecases"`
	Description        string    `json:"description"`
	EmbeddedLink       bool      `json:"embeddedLink"`
	EmbeddedPhone      bool      `json:"embeddedPhone"`
	AffiliateMarketing bool      `json:"affiliateMarketing"`
	NumberPool         bool      `json:"numberPool"`
	AgeGated           bool      `json:"ageGated"`
	DirectLending      bool      `json:"directLending"`
	SubscriberOptin    bool      `json:"subscriberOptin"`
	SubscriberOptout   bool      `json:"subscriberOptout"`
	SubscriberHelp     bool      `json:"subscriberHelp"`
	Sample1            string    `json:"sample1"`
	Sample2            string    `json:"sample2"`
	Sample3            string    `json:"sample3"`
	Sample4            string    `json:"sample4"`
	Sample5            string    `json:"sample5"`
	MessageFlow        string    `json:"messageFlow"`
	HelpKeywords       string    `json:"helpKeywords"`
	HelpMessage        string    `json:"helpMessage"`
	OptinKeywords      string    `json:"optinKeywords"`
	OptinMessage       string    `json:"optinMessage"`
	OptoutKeywords     string    `json:"optoutKeywords"`
	OptoutMessage      string    `json:"optoutMessage"`
}

// CreateBrand submits a brand for registration
func (t *textGrid) CreateCampaign(payload CreateCampaignPayload) (*Campaign, error) {
	result := new(Campaign)

	if err := t.post("campaigns/campaign", payload, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (t *textGrid) DeactivateCampaign(id string) error {
	return t.delete("campaigns/campaign/"+id, nil)
}