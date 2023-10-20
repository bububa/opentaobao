package wlbimports

import (
	"sync"
)

// AppointmentCreateResponse 结构体
type AppointmentCreateResponse struct {
	// 预约单code
	HandoverOrderCode string `json:"handover_order_code,omitempty" xml:"handover_order_code,omitempty"`
	// 预约单id
	HandoverOrderId int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
}

var poolAppointmentCreateResponse = sync.Pool{
	New: func() any {
		return new(AppointmentCreateResponse)
	},
}

// GetAppointmentCreateResponse() 从对象池中获取AppointmentCreateResponse
func GetAppointmentCreateResponse() *AppointmentCreateResponse {
	return poolAppointmentCreateResponse.Get().(*AppointmentCreateResponse)
}

// ReleaseAppointmentCreateResponse 释放AppointmentCreateResponse
func ReleaseAppointmentCreateResponse(v *AppointmentCreateResponse) {
	v.HandoverOrderCode = ""
	v.HandoverOrderId = 0
	poolAppointmentCreateResponse.Put(v)
}
