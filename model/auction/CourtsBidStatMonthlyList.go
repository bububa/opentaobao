package auction

import (
	"sync"
)

// CourtsBidStatMonthlyList 结构体
type CourtsBidStatMonthlyList struct {
	// 时间区间(月份)
	Period string `json:"period,omitempty" xml:"period,omitempty"`
	// 法院按月维度标的统计汇总
	CourtsBidStatSum *CourtsBidStatSum `json:"courts_bid_stat_sum,omitempty" xml:"courts_bid_stat_sum,omitempty"`
}

var poolCourtsBidStatMonthlyList = sync.Pool{
	New: func() any {
		return new(CourtsBidStatMonthlyList)
	},
}

// GetCourtsBidStatMonthlyList() 从对象池中获取CourtsBidStatMonthlyList
func GetCourtsBidStatMonthlyList() *CourtsBidStatMonthlyList {
	return poolCourtsBidStatMonthlyList.Get().(*CourtsBidStatMonthlyList)
}

// ReleaseCourtsBidStatMonthlyList 释放CourtsBidStatMonthlyList
func ReleaseCourtsBidStatMonthlyList(v *CourtsBidStatMonthlyList) {
	v.Period = ""
	v.CourtsBidStatSum = nil
	poolCourtsBidStatMonthlyList.Put(v)
}
