package alsc

import (
	"sync"
)

// PrizeInfo 结构体
type PrizeInfo struct {
	// 优惠券列表
	PrizeItemInfoList []PrizeItemInfo `json:"prize_item_info_list,omitempty" xml:"prize_item_info_list>prize_item_info,omitempty"`
	// 是否中奖
	WinPrize bool `json:"win_prize,omitempty" xml:"win_prize,omitempty"`
}

var poolPrizeInfo = sync.Pool{
	New: func() any {
		return new(PrizeInfo)
	},
}

// GetPrizeInfo() 从对象池中获取PrizeInfo
func GetPrizeInfo() *PrizeInfo {
	return poolPrizeInfo.Get().(*PrizeInfo)
}

// ReleasePrizeInfo 释放PrizeInfo
func ReleasePrizeInfo(v *PrizeInfo) {
	v.PrizeItemInfoList = v.PrizeItemInfoList[:0]
	v.WinPrize = false
	poolPrizeInfo.Put(v)
}
