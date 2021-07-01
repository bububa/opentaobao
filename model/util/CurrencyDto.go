package util

// CurrencyDto 结构体
type CurrencyDto struct {
	// 货币编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 货币显示标示
	Symbol string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// 货币名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
