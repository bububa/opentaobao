package aedropshiper

import (
	"sync"
)

// Money 结构体
type Money struct {
	// currencyCode
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// amount
	Amount float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// cent
	Cent int64 `json:"cent,omitempty" xml:"cent,omitempty"`
}

var poolMoney = sync.Pool{
	New: func() any {
		return new(Money)
	},
}

// GetMoney() 从对象池中获取Money
func GetMoney() *Money {
	return poolMoney.Get().(*Money)
}

// ReleaseMoney 释放Money
func ReleaseMoney(v *Money) {
	v.CurrencyCode = ""
	v.Amount = 0
	v.Cent = 0
	poolMoney.Put(v)
}
