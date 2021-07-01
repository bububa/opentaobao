package aedropshiper

// SimpleMoney 结构体
type SimpleMoney struct {
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
}
