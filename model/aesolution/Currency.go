package aesolution

import (
	"sync"
)

// Currency 结构体
type Currency struct {
	// Currency code
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// Symbol of the currency
	Symbol string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// Display name of the currency
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// The default minimum accuracy of the currency
	DefaultFractionDigits int64 `json:"default_fraction_digits,omitempty" xml:"default_fraction_digits,omitempty"`
	// Numeric code of the currency
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
	v.CurrencyCode = ""
	v.Symbol = ""
	v.DisplayName = ""
	v.DefaultFractionDigits = 0
	v.NumericCode = 0
	poolCurrency.Put(v)
}
