package examination

import (
	"sync"
)

// OrderInfo 结构体
type OrderInfo struct {
	// 外部订单ID
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 聊天地址，用户之间进行隔离
	ImUrl string `json:"im_url,omitempty" xml:"im_url,omitempty"`
	// 医生ID
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderInfo = sync.Pool{
	New: func() any {
		return new(OrderInfo)
	},
}

// GetOrderInfo() 从对象池中获取OrderInfo
func GetOrderInfo() *OrderInfo {
	return poolOrderInfo.Get().(*OrderInfo)
}

// ReleaseOrderInfo 释放OrderInfo
func ReleaseOrderInfo(v *OrderInfo) {
	v.OuterOrderId = ""
	v.Status = ""
	v.ImUrl = ""
	v.DoctorId = ""
	v.OrderId = ""
	poolOrderInfo.Put(v)
}
