package auction

import (
	"sync"
)

// CourtsBidStatSum 结构体
type CourtsBidStatSum struct {
	// 发拍次数
	PublishCount int64 `json:"publish_count,omitempty" xml:"publish_count,omitempty"`
	// 成交金额
	HammerPrice int64 `json:"hammer_price,omitempty" xml:"hammer_price,omitempty"`
	// 成交件数
	HammerCount int64 `json:"hammer_count,omitempty" xml:"hammer_count,omitempty"`
	// 结束网拍次数
	EndCount int64 `json:"end_count,omitempty" xml:"end_count,omitempty"`
	// 结束标的件数（去重）
	EndCountDist int64 `json:"end_count_dist,omitempty" xml:"end_count_dist,omitempty"`
	// 平均成交溢价率(万分位)
	AvgAddvPercent int64 `json:"avg_addv_percent,omitempty" xml:"avg_addv_percent,omitempty"`
	// 发拍件数（去重）
	PublishCountDist int64 `json:"publish_count_dist,omitempty" xml:"publish_count_dist,omitempty"`
	// 开拍件数（去重）
	StartCountDist int64 `json:"start_count_dist,omitempty" xml:"start_count_dist,omitempty"`
	// 出价次数
	BidCount int64 `json:"bid_count,omitempty" xml:"bid_count,omitempty"`
	// 起拍价（成交标的）
	StartPrice int64 `json:"start_price,omitempty" xml:"start_price,omitempty"`
	// 意向用户数（交保数+订阅数）
	InterestCount int64 `json:"interest_count,omitempty" xml:"interest_count,omitempty"`
	// 开拍数
	StartCount int64 `json:"start_count,omitempty" xml:"start_count,omitempty"`
	// 报名人数（含交保失败）
	ApplyCount int64 `json:"apply_count,omitempty" xml:"apply_count,omitempty"`
}

var poolCourtsBidStatSum = sync.Pool{
	New: func() any {
		return new(CourtsBidStatSum)
	},
}

// GetCourtsBidStatSum() 从对象池中获取CourtsBidStatSum
func GetCourtsBidStatSum() *CourtsBidStatSum {
	return poolCourtsBidStatSum.Get().(*CourtsBidStatSum)
}

// ReleaseCourtsBidStatSum 释放CourtsBidStatSum
func ReleaseCourtsBidStatSum(v *CourtsBidStatSum) {
	v.PublishCount = 0
	v.HammerPrice = 0
	v.HammerCount = 0
	v.EndCount = 0
	v.EndCountDist = 0
	v.AvgAddvPercent = 0
	v.PublishCountDist = 0
	v.StartCountDist = 0
	v.BidCount = 0
	v.StartPrice = 0
	v.InterestCount = 0
	v.StartCount = 0
	v.ApplyCount = 0
	poolCourtsBidStatSum.Put(v)
}
