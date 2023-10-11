package wlbimports

// AppointmentOrderStatusRequest 结构体
type AppointmentOrderStatusRequest struct {
	// 预约单id
	AppointmentOrderId int64 `json:"appointment_order_id,omitempty" xml:"appointment_order_id,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页面下标，从1开始
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
