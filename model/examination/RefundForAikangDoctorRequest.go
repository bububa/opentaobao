package examination

// RefundForAikangDoctorRequest 结构体
type RefundForAikangDoctorRequest struct {
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 医生id
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
}
