package icbulogistics

import (
	"sync"
)

// Money 结构体
type Money struct {
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
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
	v.Amount = ""
	v.Currency = ""
	poolMoney.Put(v)
}
