package logistic

// Money 结构体
type Money struct {
	// 币种三字码
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 分
	Cent int64 `json:"cent,omitempty" xml:"cent,omitempty"`
}
