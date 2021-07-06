package maitix

// LockTicketResponse 结构体
type LockTicketResponse struct {
	// 子订单列表
	SubOrderDtos []LockTicketSubOrderDto `json:"sub_order_dtos,omitempty" xml:"sub_order_dtos>lock_ticket_sub_order_dto,omitempty"`
	// 大麦订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 总价
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 快递费
	ExpressFee int64 `json:"express_fee,omitempty" xml:"express_fee,omitempty"`
}
