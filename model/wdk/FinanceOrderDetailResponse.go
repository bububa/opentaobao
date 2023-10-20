package wdk

import (
	"sync"
)

// FinanceOrderDetailResponse 结构体
type FinanceOrderDetailResponse struct {
	// 财务订单信息
	FinanceOrderDetails []FinanceOrderDetail `json:"finance_order_details,omitempty" xml:"finance_order_details>finance_order_detail,omitempty"`
	// 分页信息
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
}

var poolFinanceOrderDetailResponse = sync.Pool{
	New: func() any {
		return new(FinanceOrderDetailResponse)
	},
}

// GetFinanceOrderDetailResponse() 从对象池中获取FinanceOrderDetailResponse
func GetFinanceOrderDetailResponse() *FinanceOrderDetailResponse {
	return poolFinanceOrderDetailResponse.Get().(*FinanceOrderDetailResponse)
}

// ReleaseFinanceOrderDetailResponse 释放FinanceOrderDetailResponse
func ReleaseFinanceOrderDetailResponse(v *FinanceOrderDetailResponse) {
	v.FinanceOrderDetails = v.FinanceOrderDetails[:0]
	v.Pagination = nil
	poolFinanceOrderDetailResponse.Put(v)
}
