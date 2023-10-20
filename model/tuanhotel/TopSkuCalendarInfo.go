package tuanhotel

import (
	"sync"
)

// TopSkuCalendarInfo 结构体
type TopSkuCalendarInfo struct {
	// 日历库存价格信息
	Diffs []TopSkuDailyInfo `json:"diffs,omitempty" xml:"diffs>top_sku_daily_info,omitempty"`
	// 日历库存价格信息，日历库存必填
	Diff []TopSkuDailyInfo `json:"diff,omitempty" xml:"diff>top_sku_daily_info,omitempty"`
	// 日历库存范围结束日期
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 日历库存范围开始日期
	Begin string `json:"begin,omitempty" xml:"begin,omitempty"`
}

var poolTopSkuCalendarInfo = sync.Pool{
	New: func() any {
		return new(TopSkuCalendarInfo)
	},
}

// GetTopSkuCalendarInfo() 从对象池中获取TopSkuCalendarInfo
func GetTopSkuCalendarInfo() *TopSkuCalendarInfo {
	return poolTopSkuCalendarInfo.Get().(*TopSkuCalendarInfo)
}

// ReleaseTopSkuCalendarInfo 释放TopSkuCalendarInfo
func ReleaseTopSkuCalendarInfo(v *TopSkuCalendarInfo) {
	v.Diffs = v.Diffs[:0]
	v.Diff = v.Diff[:0]
	v.End = ""
	v.Begin = ""
	poolTopSkuCalendarInfo.Put(v)
}
