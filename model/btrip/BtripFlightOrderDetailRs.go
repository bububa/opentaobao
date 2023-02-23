package btrip

// BtripFlightOrderDetailRs 结构体
type BtripFlightOrderDetailRs struct {
	// 航班信息列表
	FlightInfoList []OrderFlightInfo `json:"flight_info_list,omitempty" xml:"flight_info_list>order_flight_info,omitempty"`
	// 票信息
	TicketInfoList []OrderTicketInfo `json:"ticket_info_list,omitempty" xml:"ticket_info_list>order_ticket_info,omitempty"`
	// 乘机人列表
	TravelerInfoList []OrderTravelerInfo `json:"traveler_info_list,omitempty" xml:"traveler_info_list>order_traveler_info,omitempty"`
	// 支付宝交易流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 预订人用户id
	BookUserId string `json:"book_user_id,omitempty" xml:"book_user_id,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 最迟支付时间
	LastPayTime string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// 外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 商旅订单号
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// 支付状态：0（初始状态），1（冻结成功），3（解冻成功），5（转交易成功），9（创建交易成功），11（关闭交易成功）
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 优惠金额
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 结算金额
	SettleAmount int64 `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	// 结算类型
	SettleType int64 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// 订单状态：0（初始状态）4（处理中）5（待支付）10（失败）32（订单可支付状态）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 总机建费用
	TotalBuildPrice int64 `json:"total_build_price,omitempty" xml:"total_build_price,omitempty"`
	// 总燃油费用
	TotalOilPrice int64 `json:"total_oil_price,omitempty" xml:"total_oil_price,omitempty"`
	// 总订单费用
	TotalOrderPrice int64 `json:"total_order_price,omitempty" xml:"total_order_price,omitempty"`
}
