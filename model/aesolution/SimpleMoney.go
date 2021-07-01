package aesolution

// SimpleMoney 结构体
type SimpleMoney struct {
	// currency code
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// amount
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}
