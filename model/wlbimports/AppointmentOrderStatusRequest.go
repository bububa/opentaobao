package wlbimports

// AppointmentOrderStatusRequest 结构体
type AppointmentOrderStatusRequest struct {
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 预约单id
	AppointmentOrderId int64 `json:"appointment_order_id,omitempty" xml:"appointment_order_id,omitempty"`
}
