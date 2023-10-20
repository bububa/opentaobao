package happytrip

import (
	"sync"
)

// PriceInfo 结构体
type PriceInfo struct {
	// 预估车费
	Estimate string `json:"estimate,omitempty" xml:"estimate,omitempty"`
}

var poolPriceInfo = sync.Pool{
	New: func() any {
		return new(PriceInfo)
	},
}

// GetPriceInfo() 从对象池中获取PriceInfo
func GetPriceInfo() *PriceInfo {
	return poolPriceInfo.Get().(*PriceInfo)
}

// ReleasePriceInfo 释放PriceInfo
func ReleasePriceInfo(v *PriceInfo) {
	v.Estimate = ""
	poolPriceInfo.Put(v)
}
