package omniorder

import (
	"sync"
)

// ExpandCardInfo 结构体
type ExpandCardInfo struct {
	// 用卡订单使用本金，用卡的订单才有输出，单位：分
	ExpandPriceUsed int64 `json:"expand_price_used,omitempty" xml:"expand_price_used,omitempty"`
	// 用卡订单使用权益金，用卡的订单才有输出，单位：分
	BasicPriceUsed int64 `json:"basic_price_used,omitempty" xml:"basic_price_used,omitempty"`
}

var poolExpandCardInfo = sync.Pool{
	New: func() any {
		return new(ExpandCardInfo)
	},
}

// GetExpandCardInfo() 从对象池中获取ExpandCardInfo
func GetExpandCardInfo() *ExpandCardInfo {
	return poolExpandCardInfo.Get().(*ExpandCardInfo)
}

// ReleaseExpandCardInfo 释放ExpandCardInfo
func ReleaseExpandCardInfo(v *ExpandCardInfo) {
	v.ExpandPriceUsed = 0
	v.BasicPriceUsed = 0
	poolExpandCardInfo.Put(v)
}
