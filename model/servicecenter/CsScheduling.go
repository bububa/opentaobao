package servicecenter

import (
	"sync"
)

// CsScheduling 结构体
type CsScheduling struct {
	// 一天内排班信息
	Schedulings []Scheduling `json:"schedulings,omitempty" xml:"schedulings>scheduling,omitempty"`
	// 排班日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 排班日期字符串
	StringDate string `json:"string_date,omitempty" xml:"string_date,omitempty"`
	// 排班记录更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolCsScheduling = sync.Pool{
	New: func() any {
		return new(CsScheduling)
	},
}

// GetCsScheduling() 从对象池中获取CsScheduling
func GetCsScheduling() *CsScheduling {
	return poolCsScheduling.Get().(*CsScheduling)
}

// ReleaseCsScheduling 释放CsScheduling
func ReleaseCsScheduling(v *CsScheduling) {
	v.Schedulings = v.Schedulings[:0]
	v.Date = ""
	v.StringDate = ""
	v.ModifiedTime = ""
	v.OrderId = 0
	poolCsScheduling.Put(v)
}
