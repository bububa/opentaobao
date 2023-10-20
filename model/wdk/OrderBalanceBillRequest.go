package wdk

import (
	"sync"
)

// OrderBalanceBillRequest 结构体
type OrderBalanceBillRequest struct {
	// 业务日期
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// storeId
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 最大的id
	MaxId int64 `json:"max_id,omitempty" xml:"max_id,omitempty"`
}

var poolOrderBalanceBillRequest = sync.Pool{
	New: func() any {
		return new(OrderBalanceBillRequest)
	},
}

// GetOrderBalanceBillRequest() 从对象池中获取OrderBalanceBillRequest
func GetOrderBalanceBillRequest() *OrderBalanceBillRequest {
	return poolOrderBalanceBillRequest.Get().(*OrderBalanceBillRequest)
}

// ReleaseOrderBalanceBillRequest 释放OrderBalanceBillRequest
func ReleaseOrderBalanceBillRequest(v *OrderBalanceBillRequest) {
	v.Thedate = ""
	v.StoreId = ""
	v.PageSize = 0
	v.PageNo = 0
	v.MaxId = 0
	poolOrderBalanceBillRequest.Put(v)
}
