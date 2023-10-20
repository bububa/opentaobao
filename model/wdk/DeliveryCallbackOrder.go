package wdk

import (
	"sync"
)

// DeliveryCallbackOrder 结构体
type DeliveryCallbackOrder struct {
	// 拒收子单列表
	RefusedOrderDetails []DeliveryCallbackOrderDetail `json:"refused_order_details,omitempty" xml:"refused_order_details>delivery_callback_order_detail,omitempty"`
	// 作业单号
	WorkOrderId string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
	// 作业状态变更类型：SHIP(&#34;揽收&#34;),SIGN(&#34;妥投&#34;),SIGN_ERROR(&#34;配送异常&#34;),REFUSE(&#34;拒收&#34;)
	StatusChangeType string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
	// 作业状态变更时间
	StatusChangeTime string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 配送站编码
	DeliveryDockCode string `json:"delivery_dock_code,omitempty" xml:"delivery_dock_code,omitempty"`
	// 来源系统:：CHINA_POST:邮政
	SourceSystem string `json:"source_system,omitempty" xml:"source_system,omitempty"`
	// 配送员
	Deliveryman *Deliveryman `json:"deliveryman,omitempty" xml:"deliveryman,omitempty"`
}

var poolDeliveryCallbackOrder = sync.Pool{
	New: func() any {
		return new(DeliveryCallbackOrder)
	},
}

// GetDeliveryCallbackOrder() 从对象池中获取DeliveryCallbackOrder
func GetDeliveryCallbackOrder() *DeliveryCallbackOrder {
	return poolDeliveryCallbackOrder.Get().(*DeliveryCallbackOrder)
}

// ReleaseDeliveryCallbackOrder 释放DeliveryCallbackOrder
func ReleaseDeliveryCallbackOrder(v *DeliveryCallbackOrder) {
	v.RefusedOrderDetails = v.RefusedOrderDetails[:0]
	v.WorkOrderId = ""
	v.StatusChangeType = ""
	v.StatusChangeTime = ""
	v.Remark = ""
	v.DeliveryDockCode = ""
	v.SourceSystem = ""
	v.Deliveryman = nil
	poolDeliveryCallbackOrder.Put(v)
}
