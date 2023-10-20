package aliexpresssumaitong

import (
	"sync"
)

// Currency 结构体
type Currency struct {
	// 交易币种编码
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
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
	poolCurrency.Put(v)
}
