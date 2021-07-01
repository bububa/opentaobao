package alihealthmedical

// RefuseOrderRequestDto 结构体
type RefuseOrderRequestDto struct {
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 拒绝原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 会话ID
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 互联网医院编码
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
	// 医生UUID
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 拒诊的触发类型：doctor或platform 。医生手动拒诊：doctor；三方系统触发：platform
	TriggerType string `json:"trigger_type,omitempty" xml:"trigger_type,omitempty"`
}
