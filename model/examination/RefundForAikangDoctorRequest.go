package examination

import (
	"sync"
)

// RefundForAikangDoctorRequest 结构体
type RefundForAikangDoctorRequest struct {
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 医生id
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
}

var poolRefundForAikangDoctorRequest = sync.Pool{
	New: func() any {
		return new(RefundForAikangDoctorRequest)
	},
}

// GetRefundForAikangDoctorRequest() 从对象池中获取RefundForAikangDoctorRequest
func GetRefundForAikangDoctorRequest() *RefundForAikangDoctorRequest {
	return poolRefundForAikangDoctorRequest.Get().(*RefundForAikangDoctorRequest)
}

// ReleaseRefundForAikangDoctorRequest 释放RefundForAikangDoctorRequest
func ReleaseRefundForAikangDoctorRequest(v *RefundForAikangDoctorRequest) {
	v.OrderId = ""
	v.DoctorId = ""
	v.RefundReason = ""
	poolRefundForAikangDoctorRequest.Put(v)
}
