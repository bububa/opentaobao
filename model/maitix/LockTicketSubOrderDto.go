package maitix

// LockTicketSubOrderDto 结构体
type LockTicketSubOrderDto struct {
	// 外部子订单号，透传返回
	ExternalSubOrderNo string `json:"external_sub_order_no,omitempty" xml:"external_sub_order_no,omitempty"`
	// 商品原价，单位分
	OriginPrice int64 `json:"origin_price,omitempty" xml:"origin_price,omitempty"`
	// 商品实际价，单位分
	RealPrice int64 `json:"real_price,omitempty" xml:"real_price,omitempty"`
	// 大麦子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 子订单座位信息
	SubOrderSeatDto *LockTicketSubOrderSeatDto `json:"sub_order_seat_dto,omitempty" xml:"sub_order_seat_dto,omitempty"`
	// 票单ID
	VoucherId int64 `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
}
