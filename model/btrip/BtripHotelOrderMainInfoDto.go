package btrip

// BtripHotelOrderMainInfoDto 结构体
type BtripHotelOrderMainInfoDto struct {
	// 入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 分销商订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 最晚到店时间，可能为空
	LateArriveTime string `json:"late_arrive_time,omitempty" xml:"late_arrive_time,omitempty"`
	// 订单创建日期
	OrderCreateDate string `json:"order_create_date,omitempty" xml:"order_create_date,omitempty"`
	// 订单状态描述
	OrderStatusDesc string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// 供应商订单id
	SupplierOrderId string `json:"supplier_order_id,omitempty" xml:"supplier_order_id,omitempty"`
	// 实际入住时间
	RealCheckinTime string `json:"real_checkin_time,omitempty" xml:"real_checkin_time,omitempty"`
	// 实际离店时间
	RealCheckoutTime string `json:"real_checkout_time,omitempty" xml:"real_checkout_time,omitempty"`
	// 商旅订单id
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// 买家实退金额
	BuyerRealRefund int64 `json:"buyer_real_refund,omitempty" xml:"buyer_real_refund,omitempty"`
	// 卖家优惠的金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 住几晚
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 平台优惠金额，单位：分
	PlatformPromotionAmt int64 `json:"platform_promotion_amt,omitempty" xml:"platform_promotion_amt,omitempty"`
	// 订单房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 实际支付金额
	TotalActualPrice int64 `json:"total_actual_price,omitempty" xml:"total_actual_price,omitempty"`
	// 总房费
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
}
