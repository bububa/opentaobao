package xhotelonlineorder

// Toprelationorders 结构体
type Toprelationorders struct {
	// 订单tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// 联系人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 联系手机号
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 统一订单状态
	UnionStatusText string `json:"union_status_text,omitempty" xml:"union_status_text,omitempty"`
	// 统一订单状态
	UnionStatusValue string `json:"union_status_value,omitempty" xml:"union_status_value,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 入住天数
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 支付金额
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
}
