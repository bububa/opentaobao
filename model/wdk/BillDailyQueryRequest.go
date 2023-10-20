package wdk

import (
	"sync"
)

// BillDailyQueryRequest 结构体
type BillDailyQueryRequest struct {
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 当前页码，从1开始
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 账单开始时间，默认系统时间前一天
	BillStartDate int64 `json:"bill_start_date,omitempty" xml:"bill_start_date,omitempty"`
	// 账单结束时间，默认系统时间前一天
	BillEndDate int64 `json:"bill_end_date,omitempty" xml:"bill_end_date,omitempty"`
}

var poolBillDailyQueryRequest = sync.Pool{
	New: func() any {
		return new(BillDailyQueryRequest)
	},
}

// GetBillDailyQueryRequest() 从对象池中获取BillDailyQueryRequest
func GetBillDailyQueryRequest() *BillDailyQueryRequest {
	return poolBillDailyQueryRequest.Get().(*BillDailyQueryRequest)
}

// ReleaseBillDailyQueryRequest 释放BillDailyQueryRequest
func ReleaseBillDailyQueryRequest(v *BillDailyQueryRequest) {
	v.StoreId = ""
	v.Current = 0
	v.PageSize = 0
	v.BillStartDate = 0
	v.BillEndDate = 0
	poolBillDailyQueryRequest.Put(v)
}
