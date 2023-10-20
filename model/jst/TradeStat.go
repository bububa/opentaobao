package jst

import (
	"sync"
)

// TradeStat 结构体
type TradeStat struct {
	// 状态名称
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolTradeStat = sync.Pool{
	New: func() any {
		return new(TradeStat)
	},
}

// GetTradeStat() 从对象池中获取TradeStat
func GetTradeStat() *TradeStat {
	return poolTradeStat.Get().(*TradeStat)
}

// ReleaseTradeStat 释放TradeStat
func ReleaseTradeStat(v *TradeStat) {
	v.Status = ""
	v.Count = 0
	poolTradeStat.Put(v)
}
