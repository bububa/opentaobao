package aedropshiper

// SimpleMoney 结构体
type SimpleMoney struct {
	// Amount
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// Currency
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
}
