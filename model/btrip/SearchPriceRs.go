package btrip

import (
	"sync"
)

// SearchPriceRs 结构体
type SearchPriceRs struct {
	// 原始销售价，活动前的价格（单位：元）
	OriginalSellPrice int64 `json:"original_sell_price,omitempty" xml:"original_sell_price,omitempty"`
	// 销售价(单位：元)
	SellPrice int64 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// 参考税(单位：元)
	Tax int64 `json:"tax,omitempty" xml:"tax,omitempty"`
}

var poolSearchPriceRs = sync.Pool{
	New: func() any {
		return new(SearchPriceRs)
	},
}

// GetSearchPriceRs() 从对象池中获取SearchPriceRs
func GetSearchPriceRs() *SearchPriceRs {
	return poolSearchPriceRs.Get().(*SearchPriceRs)
}

// ReleaseSearchPriceRs 释放SearchPriceRs
func ReleaseSearchPriceRs(v *SearchPriceRs) {
	v.OriginalSellPrice = 0
	v.SellPrice = 0
	v.Tax = 0
	poolSearchPriceRs.Put(v)
}
