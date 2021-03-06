package alihealthmedical

// OrderQueryRequestDto 结构体
type OrderQueryRequestDto struct {
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 互联网医院编码
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}
