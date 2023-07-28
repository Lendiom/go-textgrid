package textgrid

type TcrWebhookPayload struct {
	EventType string `json:"eventType"`

	CspID   string `json:"cspId"`
	CspName string `json:"cspName"`

	BrandName           string `json:"brandName"`
	BrandReferenceID    any    `json:"brandReferenceId"`
	BrandID             string `json:"brandId"`
	BrandIdentityStatus string `json:"brandIdentityStatus"`

	CampaignID          string `json:"campaignId"`
	CampaignReferenceID string `json:"campaignReferenceId"`

	Description string `json:"description"`
	Mock        bool   `json:"mock"`
}

type TcrWebhookEventType string

var (
	TcrWebhookEventTypeBrandIdentityStatusUpdate TcrWebhookEventType = "BRAND_IDENTITY_STATUS_UPDATE"
	TcrWebhookEventTypeBrandDelete               TcrWebhookEventType = "BRAND_DELETE"
	TcrWebhookEventTypeCampaignDCAComplete       TcrWebhookEventType = "CAMPAIGN_DCA_COMPLETE"
	TcrWebhookEventTypeCampaignBilled            TcrWebhookEventType = "CAMPAIGN_BILLED"
	TcrWebhookEventTypeCampaignExpired           TcrWebhookEventType = "CAMPAIGN_EXPIRED"
)
