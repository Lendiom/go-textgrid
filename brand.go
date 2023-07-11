package textgrid

type CreateBrandPayload struct {
	BrandID string `json:"brandId"`
	CspID   string `json:"cspId"`

	EntityType        BrandEntityType `json:"entityType"`
	FirstName         string          `json:"firstName,omitempty"`
	LastName          string          `json:"lastName,omitempty"`
	DisplayName       string          `json:"displayName,omitempty"`
	CompanyName       string          `json:"companyName,omitempty"`
	Ein               string          `json:"ein,omitempty"`
	EinIssuingCountry string          `json:"einIssuingCountry,omitempty"`
	Phone             string          `json:"phone,omitempty"`
	MobilePhone       string          `json:"mobilePhone,omitempty"`
	Street            string          `json:"street,omitempty"`
	City              string          `json:"city,omitempty"`
	State             string          `json:"state,omitempty"`
	PostalCode        string          `json:"postalCode,omitempty"`
	Country           string          `json:"country,omitempty"`
	Email             string          `json:"email,omitempty"`
	StockSymbol       string          `json:"stockSymbol,omitempty"`
	StockExchange     string          `json:"stockExchange,omitempty"` //TODO: use the enumerated values
	IPAddress         string          `json:"ipAddress,omitempty"`
	Website           string          `json:"website,omitempty"`
	BrandRelationship string          `json:"brandRelationship,omitempty"` //TODO: use the enumerated values
	Vertical          string          `json:"vertical,omitempty"`          //TODO: use the enumerated values
	AltBusinessID     string          `json:"altBusinessId,omitempty"`
	AltBusinessIDType string          `json:"altBusinessIdType,omitempty"` //TODO: use the enumerated values
	ReferenceID       string          `json:"referenceId,omitempty"`
	Mock              bool            `json:"mock"`
}

type Brand struct {
	EntityType        BrandEntityType `json:"entityType"`
	CspID             string          `json:"cspId"`
	BrandID           string          `json:"brandId"`
	FirstName         string          `json:"firstName"`
	LastName          string          `json:"lastName"`
	DisplayName       string          `json:"displayName"`
	CompanyName       string          `json:"companyName"`
	Ein               string          `json:"ein"`
	EinIssuingCountry string          `json:"einIssuingCountry"`
	Phone             string          `json:"phone"`
	MobilePhone       string          `json:"mobilePhone"`
	Street            string          `json:"street"`
	City              string          `json:"city"`
	State             string          `json:"state"`
	PostalCode        string          `json:"postalCode"`
	Country           string          `json:"country"`
	Email             string          `json:"email"`
	StockSymbol       string          `json:"stockSymbol"`
	StockExchange     string          `json:"stockExchange"` //TODO: use the enumerated values
	IPAddress         string          `json:"ipAddress"`
	Website           string          `json:"website"`
	BrandRelationship string          `json:"brandRelationship"` //TODO: use the enumerated values
	Vertical          string          `json:"vertical"`          //TODO: use the enumerated values
	AltBusinessID     string          `json:"altBusinessId"`
	AltBusinessIDType string          `json:"altBusinessIdType"` //TODO: use the enumerated values
	UniversalEin      string          `json:"universalEin"`
	ReferenceID       string          `json:"referenceId"`
	Mock              bool            `json:"mock"`
	IdentityStatus    string          `json:"identityStatus"` //TODO: is there an enum for this?
}

type BrandEntityType string

var (
	BrandEntityTypePrivateProfit  BrandEntityType = "PRIVATE_PROFIT"
	BrandEntityTypePublicProfit   BrandEntityType = "PUBLIC_PROFIT"
	BrandEntityTypeNonProfit      BrandEntityType = "NON_PROFIT"
	BrandEntityTypeGovernment     BrandEntityType = "GOVERNMENT"
	BrandEntityTypeSoleProprietor BrandEntityType = "SOLE_PROPRIETOR"
)

// CreateBrand submits a brand for registration
func (t *textGrid) CreateBrand(brand CreateBrandPayload) (*Brand, error) {
	result := new(Brand)

	if err := t.post("campaigns/brand", brand, result); err != nil {
		return nil, err
	}

	return result, nil
}
