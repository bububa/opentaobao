package xhotel

import (
	"sync"
)

// LongOrderInfo 结构体
type LongOrderInfo struct {
	// 最小连住天数
	MinContinuityStay int64 `json:"min_continuity_stay,omitempty" xml:"min_continuity_stay,omitempty"`
	// 折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
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
	v.MinContinuityStay = 0
	v.InvestmentNumber = 0
	poolLongOrderInfo.Put(v)
}
