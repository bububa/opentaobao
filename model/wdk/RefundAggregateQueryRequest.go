package wdk

import (
	"sync"
)

// RefundAggregateQueryRequest 结构体
type RefundAggregateQueryRequest struct {
	// 起始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 下单终端: APP / POS
	OrderClient string `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 收营员id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 分页序号
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页size
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolRefundAggregateQueryRequest = sync.Pool{
	New: func() any {
		return new(RefundAggregateQueryRequest)
	},
}

// GetRefundAggregateQueryRequest() 从对象池中获取RefundAggregateQueryRequest
func GetRefundAggregateQueryRequest() *RefundAggregateQueryRequest {
	return poolRefundAggregateQueryRequest.Get().(*RefundAggregateQueryRequest)
}

// ReleaseRefundAggregateQueryRequest 释放RefundAggregateQueryRequest
func ReleaseRefundAggregateQueryRequest(v *RefundAggregateQueryRequest) {
	v.StartTime = ""
	v.EndTime = ""
	v.StoreId = ""
	v.OrderClient = ""
	v.OperatorId = ""
	v.PageIndex = 0
	v.PageSize = 0
	poolRefundAggregateQueryRequest.Put(v)
}
