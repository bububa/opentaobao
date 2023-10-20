package perfect

import (
	"sync"
)

// PickOrderRequest 结构体
type PickOrderRequest struct {
	// 拣货子单
	PickOrderDetails []PickOrderDetailRequest `json:"pick_order_details,omitempty" xml:"pick_order_details>pick_order_detail_request,omitempty"`
	// 批次号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 拣货单
	PickOrderCode string `json:"pick_order_code,omitempty" xml:"pick_order_code,omitempty"`
}

var poolPickOrderRequest = sync.Pool{
	New: func() any {
		return new(PickOrderRequest)
	},
}

// GetPickOrderRequest() 从对象池中获取PickOrderRequest
func GetPickOrderRequest() *PickOrderRequest {
	return poolPickOrderRequest.Get().(*PickOrderRequest)
}

// ReleasePickOrderRequest 释放PickOrderRequest
func ReleasePickOrderRequest(v *PickOrderRequest) {
	v.PickOrderDetails = v.PickOrderDetails[:0]
	v.BatchCode = ""
	v.Attributes = ""
	v.PickOrderCode = ""
	poolPickOrderRequest.Put(v)
}
