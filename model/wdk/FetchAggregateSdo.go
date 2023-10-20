package wdk

import (
	"sync"
)

// FetchAggregateSdo 结构体
type FetchAggregateSdo struct {
	// 主单号
	MainOutOrderId string `json:"main_out_order_id,omitempty" xml:"main_out_order_id,omitempty"`
	// 子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
	// 期望取货数量
	ExpectFetchQuantity string `json:"expect_fetch_quantity,omitempty" xml:"expect_fetch_quantity,omitempty"`
	// 时间取货数量
	ActualFetchQuantity string `json:"actual_fetch_quantity,omitempty" xml:"actual_fetch_quantity,omitempty"`
	// 取货结束时间
	FetchEndTime string `json:"fetch_end_time,omitempty" xml:"fetch_end_time,omitempty"`
	// 期望退款数量
	ExpectRefundQuantity string `json:"expect_refund_quantity,omitempty" xml:"expect_refund_quantity,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 取货类型
	FetchType int64 `json:"fetch_type,omitempty" xml:"fetch_type,omitempty"`
}

var poolFetchAggregateSdo = sync.Pool{
	New: func() any {
		return new(FetchAggregateSdo)
	},
}

// GetFetchAggregateSdo() 从对象池中获取FetchAggregateSdo
func GetFetchAggregateSdo() *FetchAggregateSdo {
	return poolFetchAggregateSdo.Get().(*FetchAggregateSdo)
}

// ReleaseFetchAggregateSdo 释放FetchAggregateSdo
func ReleaseFetchAggregateSdo(v *FetchAggregateSdo) {
	v.MainOutOrderId = ""
	v.SubOutOrderId = ""
	v.ExpectFetchQuantity = ""
	v.ActualFetchQuantity = ""
	v.FetchEndTime = ""
	v.ExpectRefundQuantity = ""
	v.Status = 0
	v.FetchType = 0
	poolFetchAggregateSdo.Put(v)
}
