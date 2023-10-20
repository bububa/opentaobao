package aesolution

import (
	"sync"
)

// SimpleMoney 结构体
type SimpleMoney struct {
	// currency code
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// amount
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolSimpleMoney = sync.Pool{
	New: func() any {
		return new(SimpleMoney)
	},
}

// GetSimpleMoney() 从对象池中获取SimpleMoney
func GetSimpleMoney() *SimpleMoney {
	return poolSimpleMoney.Get().(*SimpleMoney)
}

// ReleaseSimpleMoney 释放SimpleMoney
func ReleaseSimpleMoney(v *SimpleMoney) {
	v.CurrencyCode = ""
	v.Amount = ""
	poolSimpleMoney.Put(v)
}
