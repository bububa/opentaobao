package wlbimports

import (
	"sync"
)

// AppointmentOrderCancelRequset 结构体
type AppointmentOrderCancelRequset struct {
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 预约单id
	AppointmentOrderId int64 `json:"appointment_order_id,omitempty" xml:"appointment_order_id,omitempty"`
}

var poolAppointmentOrderCancelRequset = sync.Pool{
	New: func() any {
		return new(AppointmentOrderCancelRequset)
	},
}

// GetAppointmentOrderCancelRequset() 从对象池中获取AppointmentOrderCancelRequset
func GetAppointmentOrderCancelRequset() *AppointmentOrderCancelRequset {
	return poolAppointmentOrderCancelRequset.Get().(*AppointmentOrderCancelRequset)
}

// ReleaseAppointmentOrderCancelRequset 释放AppointmentOrderCancelRequset
func ReleaseAppointmentOrderCancelRequset(v *AppointmentOrderCancelRequset) {
	v.SellerId = 0
	v.AppointmentOrderId = 0
	poolAppointmentOrderCancelRequset.Put(v)
}
