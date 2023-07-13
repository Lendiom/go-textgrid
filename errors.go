package textgrid

type genericError struct {
	Error string `json:"Error"`
}

type badRequestError struct {
	Code        int      `json:"code"`
	Description string   `json:"description"`
	Field       string   `json:"field"`
	Fields      []string `json:"fields"`
}
