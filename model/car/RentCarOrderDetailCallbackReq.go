package car

import (
	"sync"
)

// RentCarOrderDetailCallbackReq 结构体
type RentCarOrderDetailCallbackReq struct {
	// 飞猪订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolRentCarOrderDetailCallbackReq = sync.Pool{
	New: func() any {
		return new(RentCarOrderDetailCallbackReq)
	},
}

// GetRentCarOrderDetailCallbackReq() 从对象池中获取RentCarOrderDetailCallbackReq
func GetRentCarOrderDetailCallbackReq() *RentCarOrderDetailCallbackReq {
	return poolRentCarOrderDetailCallbackReq.Get().(*RentCarOrderDetailCallbackReq)
}

// ReleaseRentCarOrderDetailCallbackReq 释放RentCarOrderDetailCallbackReq
func ReleaseRentCarOrderDetailCallbackReq(v *RentCarOrderDetailCallbackReq) {
	v.OrderId = 0
	poolRentCarOrderDetailCallbackReq.Put(v)
}
