package txcs

import (
	"sync"
)

// Currency 结构体
type Currency struct {
	// 符号
	Symbol string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// 显示中文
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// 显示英文
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 默认位数
	DefaultFractionDigits int64 `json:"default_fraction_digits,omitempty" xml:"default_fraction_digits,omitempty"`
	// 货币编码
	NumericCode int64 `json:"numeric_code,omitempty" xml:"numeric_code,omitempty"`
}

var poolCurrency = sync.Pool{
	New: func() any {
		return new(Currency)
	},
}

// GetCurrency() 从对象池中获取Currency
func GetCurrency() *Currency {
	return poolCurrency.Get().(*Currency)
}

// ReleaseCurrency 释放Currency
func ReleaseCurrency(v *Currency) {
	v.Symbol = ""
	v.DisplayName = ""
	v.CurrencyCode = ""
	v.DefaultFractionDigits = 0
	v.NumericCode = 0
	poolCurrency.Put(v)
}
