package scbp

// TopCustomerDto 结构体
type TopCustomerDto struct {
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 区域
	Region string `json:"region,omitempty" xml:"region,omitempty"`
	// 客户等级
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 分数
	Score int64 `json:"score,omitempty" xml:"score,omitempty"`
	// 下月星级
	StarTemp int64 `json:"star_temp,omitempty" xml:"star_temp,omitempty"`
	// 下月分数
	ScoreTemp int64 `json:"score_temp,omitempty" xml:"score_temp,omitempty"`
	// 账户余额，单位：分
	CustomerBalance int64 `json:"customer_balance,omitempty" xml:"customer_balance,omitempty"`
}
