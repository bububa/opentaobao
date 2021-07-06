package aliexpresssumaitong

// HjTaxCalculateResultDto 结构体
type HjTaxCalculateResultDto struct {
	// 交易行计税结果
	Lines []Lines `json:"lines,omitempty" xml:"lines>lines,omitempty"`
	// 币种转欧元汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 交易币种
	Currency *Currency `json:"currency,omitempty" xml:"currency,omitempty"`
}
