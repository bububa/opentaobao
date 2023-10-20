package maitix

import (
	"sync"
)

// Money 结构体
type Money struct {
	// 票价 单位：分
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
	v.Cent = 0
	poolMoney.Put(v)
}
