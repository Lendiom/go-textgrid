package textgrid

type CreateBrandPayload struct {
	EntityType        BrandEntityType        `json:"entityType"`
	FirstName         string                 `json:"firstName,omitempty"`
	LastName          string                 `json:"lastName,omitempty"`
	DisplayName       string                 `json:"displayName,omitempty"`
	CompanyName       string                 `json:"companyName,omitempty"`
	Ein               string                 `json:"ein,omitempty"`
	EinIssuingCountry string                 `json:"einIssuingCountry,omitempty"`
	Phone             string                 `json:"phone,omitempty"`
	MobilePhone       string                 `json:"mobilePhone,omitempty"`
	Street            string                 `json:"street,omitempty"`
	City              string                 `json:"city,omitempty"`
	State             string                 `json:"state,omitempty"`
	PostalCode        string                 `json:"postalCode,omitempty"`
	Country           string                 `json:"country,omitempty"`
	Email             string                 `json:"email,omitempty"`
	StockSymbol       string                 `json:"stockSymbol,omitempty"`
	StockExchange     string                 `json:"stockExchange,omitempty"` //TODO: use the enumerated values
	IPAddress         string                 `json:"ipAddress,omitempty"`
	Website           string                 `json:"website,omitempty"`
	BrandRelationship BrandRelationship      `json:"brandRelationship,omitempty"`
	Vertical          BrandVertical          `json:"vertical,omitempty"`
	AltBusinessID     string                 `json:"altBusinessId,omitempty"`
	AltBusinessIDType BrandAltBusinessIDType `json:"altBusinessIdType,omitempty"`
	ReferenceID       string                 `json:"referenceId,omitempty"`
	Mock              bool                   `json:"mock"`
}

type Brand struct {
	EntityType        BrandEntityType        `json:"entityType"`
	CspID             string                 `json:"cspId"`
	BrandID           string                 `json:"brandId"`
	FirstName         string                 `json:"firstName"`
	LastName          string                 `json:"lastName"`
	DisplayName       string                 `json:"displayName"`
	CompanyName       string                 `json:"companyName"`
	Ein               string                 `json:"ein"`
	EinIssuingCountry string                 `json:"einIssuingCountry"`
	Phone             string                 `json:"phone"`
	MobilePhone       string                 `json:"mobilePhone"`
	Street            string                 `json:"street"`
	City              string                 `json:"city"`
	State             string                 `json:"state"`
	PostalCode        string                 `json:"postalCode"`
	Country           string                 `json:"country"`
	Email             string                 `json:"email"`
	StockSymbol       string                 `json:"stockSymbol"`
	StockExchange     string                 `json:"stockExchange"` //TODO: use the enumerated values
	IPAddress         string                 `json:"ipAddress"`
	Website           string                 `json:"website"`
	BrandRelationship BrandRelationship      `json:"brandRelationship"`
	Vertical          BrandVertical          `json:"vertical"`
	AltBusinessID     string                 `json:"altBusinessId"`
	AltBusinessIDType BrandAltBusinessIDType `json:"altBusinessIdType"`
	UniversalEin      string                 `json:"universalEin"`
	ReferenceID       string                 `json:"referenceId"`
	Mock              bool                   `json:"mock"`
	IdentityStatus    string                 `json:"identityStatus"` //TODO: is there an enum for this?
}

type BrandEntityType string

var (
	BrandEntityTypePrivateProfit  BrandEntityType = "PRIVATE_PROFIT"
	BrandEntityTypePublicProfit   BrandEntityType = "PUBLIC_PROFIT"
	BrandEntityTypeNonProfit      BrandEntityType = "NON_PROFIT"
	BrandEntityTypeGovernment     BrandEntityType = "GOVERNMENT"
	BrandEntityTypeSoleProprietor BrandEntityType = "SOLE_PROPRIETOR"
)

type BrandRelationship string

var (
	BrandRelationshipBasicAccount  BrandRelationship = "BASIC_ACCOUNT"
	BrandRelationshipSmallAccount  BrandRelationship = "SMALL_ACCOUNT"
	BrandRelationshipMediumAccount BrandRelationship = "MEDIUM_ACCOUNT"
	BrandRelationshipLargeAccount  BrandRelationship = "LARGE_ACCOUNT"
	BrandRelationshipKeyAccount    BrandRelationship = "KEY_ACCOUNT"
)

type BrandVertical string

var (
	BrandVerticalProfessional   BrandVertical = "PROFESSIONAL"
	BrandVerticalRealEstate     BrandVertical = "REAL_ESTATE"
	BrandVerticalHealthcare     BrandVertical = "HEALTHCARE"
	BrandVerticalHumanResources BrandVertical = "HUMAN_RESOURCES"
	BrandVerticalEnergy         BrandVertical = "ENERGY"
	BrandVerticalEntertainment  BrandVertical = "ENTERTAINMENT"
	BrandVerticalRetail         BrandVertical = "RETAIL"
	BrandVerticalTransportation BrandVertical = "TRANSPORTATION"
	BrandVerticalAgriculture    BrandVertical = "AGRICULTURE"
	BrandVerticalInsurance      BrandVertical = "INSURANCE"
	BrandVerticalPostal         BrandVertical = "POSTAL"
	BrandVerticalEducation      BrandVertical = "EDUCATION"
	BrandVerticalHospitality    BrandVertical = "HOSPITALITY"
	BrandVerticalFinancial      BrandVertical = "FINANCIAL"
	BrandVerticalPolitical      BrandVertical = "POLITICAL"
	BrandVerticalGambling       BrandVertical = "GAMBLING"
	BrandVerticalLegal          BrandVertical = "LEGAL"
	BrandVerticalConstruction   BrandVertical = "CONSTRUCTION"
	BrandVerticalNonProfit      BrandVertical = "NGO"
	BrandVerticalManufacturing  BrandVertical = "MANUFACTURING"
	BrandVerticalGovernment     BrandVertical = "GOVERNMENT"
	BrandVerticalTechnology     BrandVertical = "TECHNOLOGY"
	BrandVerticalCommunication  BrandVertical = "COMMUNICATION"
)

type BrandAltBusinessIDType string

var (
	BrandAltBusinessIDTypeNone BrandAltBusinessIDType = "NONE"
	BrandAltBusinessIDTypeDUNS BrandAltBusinessIDType = "DUNS"
	BrandAltBusinessIDTypeGIIN BrandAltBusinessIDType = "GIIN"
	BrandAltBusinessIDTypeLEI  BrandAltBusinessIDType = "LEI"
)

// CreateBrand submits a brand for registration
func (t *textGrid) CreateBrand(brand CreateBrandPayload) (*Brand, error) {
	result := new(Brand)

	if err := t.post("campaigns/brand", brand, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetBrand gets the current brand information
func (t *textGrid) GetBrand(id string) (*Brand, error) {
	resp := new(Brand)

	if err := t.get("campaigns/brand/"+id, nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteBrand removes the brand, unless it has an active campaign
func (t *textGrid) DeleteBrand(id string) error {
	return t.delete("campaigns/brand/"+id, nil)
}
