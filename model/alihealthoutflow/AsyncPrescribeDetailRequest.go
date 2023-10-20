package alihealthoutflow

import (
	"sync"
)

// AsyncPrescribeDetailRequest 结构体
type AsyncPrescribeDetailRequest struct {
	// 处方号
	RxNo string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
	// 医院id
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}

var poolAsyncPrescribeDetailRequest = sync.Pool{
	New: func() any {
		return new(AsyncPrescribeDetailRequest)
	},
}

// GetAsyncPrescribeDetailRequest() 从对象池中获取AsyncPrescribeDetailRequest
func GetAsyncPrescribeDetailRequest() *AsyncPrescribeDetailRequest {
	return poolAsyncPrescribeDetailRequest.Get().(*AsyncPrescribeDetailRequest)
}

// ReleaseAsyncPrescribeDetailRequest 释放AsyncPrescribeDetailRequest
func ReleaseAsyncPrescribeDetailRequest(v *AsyncPrescribeDetailRequest) {
	v.RxNo = ""
	v.HospitalId = ""
	poolAsyncPrescribeDetailRequest.Put(v)
}
