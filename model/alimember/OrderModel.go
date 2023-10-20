package alimember

import (
	"sync"
)

// OrderModel 结构体
type OrderModel struct {
	// 单据类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 单据号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 该单据对应付费会员开始时间
	OrderIdentityEndTime string `json:"order_identity_end_time,omitempty" xml:"order_identity_end_time,omitempty"`
	// 该单据对应付费会员结束时间
	OrderIdentityStartTime string `json:"order_identity_start_time,omitempty" xml:"order_identity_start_time,omitempty"`
}

var poolOrderModel = sync.Pool{
	New: func() any {
		return new(OrderModel)
	},
}

// GetOrderModel() 从对象池中获取OrderModel
func GetOrderModel() *OrderModel {
	return poolOrderModel.Get().(*OrderModel)
}

// ReleaseOrderModel 释放OrderModel
func ReleaseOrderModel(v *OrderModel) {
	v.OrderType = ""
	v.OrderNo = ""
	v.OrderIdentityEndTime = ""
	v.OrderIdentityStartTime = ""
	poolOrderModel.Put(v)
}
