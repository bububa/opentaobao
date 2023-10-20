package xhotelitem

import (
	"sync"
)

// LongOrderInfo 结构体
type LongOrderInfo struct {
	// 互动折扣
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
	// 最小连住天数
	MinContinuityStay int64 `json:"min_continuity_stay,omitempty" xml:"min_continuity_stay,omitempty"`
}

var poolLongOrderInfo = sync.Pool{
	New: func() any {
		return new(LongOrderInfo)
	},
}

// GetLongOrderInfo() 从对象池中获取LongOrderInfo
func GetLongOrderInfo() *LongOrderInfo {
	return poolLongOrderInfo.Get().(*LongOrderInfo)
}

// ReleaseLongOrderInfo 释放LongOrderInfo
func ReleaseLongOrderInfo(v *LongOrderInfo) {
	v.InvestmentNumber = 0
	v.MinContinuityStay = 0
	poolLongOrderInfo.Put(v)
}
