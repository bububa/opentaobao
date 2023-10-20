package wdk

import (
	"sync"
)

// DeliveryCallbackOrderDetail 结构体
type DeliveryCallbackOrderDetail struct {
	// 子单ID
	WorkOrderDetailId string `json:"work_order_detail_id,omitempty" xml:"work_order_detail_id,omitempty"`
	// 拒收原因
	RefusedReason string `json:"refused_reason,omitempty" xml:"refused_reason,omitempty"`
}

var poolDeliveryCallbackOrderDetail = sync.Pool{
	New: func() any {
		return new(DeliveryCallbackOrderDetail)
	},
}

// GetDeliveryCallbackOrderDetail() 从对象池中获取DeliveryCallbackOrderDetail
func GetDeliveryCallbackOrderDetail() *DeliveryCallbackOrderDetail {
	return poolDeliveryCallbackOrderDetail.Get().(*DeliveryCallbackOrderDetail)
}

// ReleaseDeliveryCallbackOrderDetail 释放DeliveryCallbackOrderDetail
func ReleaseDeliveryCallbackOrderDetail(v *DeliveryCallbackOrderDetail) {
	v.WorkOrderDetailId = ""
	v.RefusedReason = ""
	poolDeliveryCallbackOrderDetail.Put(v)
}
