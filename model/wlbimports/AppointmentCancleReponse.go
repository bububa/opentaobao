package wlbimports

import (
	"sync"
)

// AppointmentCancleReponse 结构体
type AppointmentCancleReponse struct {
	// 取消是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolAppointmentCancleReponse = sync.Pool{
	New: func() any {
		return new(AppointmentCancleReponse)
	},
}

// GetAppointmentCancleReponse() 从对象池中获取AppointmentCancleReponse
func GetAppointmentCancleReponse() *AppointmentCancleReponse {
	return poolAppointmentCancleReponse.Get().(*AppointmentCancleReponse)
}

// ReleaseAppointmentCancleReponse 释放AppointmentCancleReponse
func ReleaseAppointmentCancleReponse(v *AppointmentCancleReponse) {
	v.Result = false
	poolAppointmentCancleReponse.Put(v)
}
