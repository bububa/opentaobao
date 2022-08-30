package wlbimports

// AppointmentCreateResponse 结构体
type AppointmentCreateResponse struct {
	// 预约单code
	HandoverOrderCode string `json:"handover_order_code,omitempty" xml:"handover_order_code,omitempty"`
	// 预约单id
	HandoverOrderId int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
}
