package wdk

import (
	"sync"
)

// FetchAggregate 结构体
type FetchAggregate struct {
	// 期望取货数量
	ExpectFetchQuantity string `json:"expect_fetch_quantity,omitempty" xml:"expect_fetch_quantity,omitempty"`
	// 期望退款数量
	ExpectRefundQuantity string `json:"expect_refund_quantity,omitempty" xml:"expect_refund_quantity,omitempty"`
	// 子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
	// 取货类型（1上门取货）
	FetchType int64 `json:"fetch_type,omitempty" xml:"fetch_type,omitempty"`
}

var poolFetchAggregate = sync.Pool{
	New: func() any {
		return new(FetchAggregate)
	},
}

// GetFetchAggregate() 从对象池中获取FetchAggregate
func GetFetchAggregate() *FetchAggregate {
	return poolFetchAggregate.Get().(*FetchAggregate)
}

// ReleaseFetchAggregate 释放FetchAggregate
func ReleaseFetchAggregate(v *FetchAggregate) {
	v.ExpectFetchQuantity = ""
	v.ExpectRefundQuantity = ""
	v.SubOutOrderId = ""
	v.FetchType = 0
	poolFetchAggregate.Put(v)
}
