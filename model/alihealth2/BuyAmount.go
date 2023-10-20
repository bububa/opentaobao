package alihealth2

import (
	"sync"
)

// BuyAmount 结构体
type BuyAmount struct {
	// 本次购买量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 统计周期，不传的话则不做周期限购
	HistoryDays int64 `json:"history_days,omitempty" xml:"history_days,omitempty"`
	// 历史购买量。历史购买量需要包含本次购买量，不传的话则不做周期限购
	HistoryTotalAmount int64 `json:"history_total_amount,omitempty" xml:"history_total_amount,omitempty"`
}

var poolBuyAmount = sync.Pool{
	New: func() any {
		return new(BuyAmount)
	},
}

// GetBuyAmount() 从对象池中获取BuyAmount
func GetBuyAmount() *BuyAmount {
	return poolBuyAmount.Get().(*BuyAmount)
}

// ReleaseBuyAmount 释放BuyAmount
func ReleaseBuyAmount(v *BuyAmount) {
	v.BuyAmount = 0
	v.HistoryDays = 0
	v.HistoryTotalAmount = 0
	poolBuyAmount.Put(v)
}
