package logistic

import (
	"sync"
)

// Money 结构体
type Money struct {
	// 币种三字码
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 分
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
	v.Cent = 0
	poolMoney.Put(v)
}
