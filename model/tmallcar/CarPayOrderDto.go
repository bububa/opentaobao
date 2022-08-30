package tmallcar

// CarPayOrderDto 结构体
type CarPayOrderDto struct {
	// 子支付单列表
	SubPayOrders []CarSubPayOrderDto `json:"sub_pay_orders,omitempty" xml:"sub_pay_orders>car_sub_pay_order_dto,omitempty"`
	// 支付单id
	PayOrderId int64 `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 支付总金额，单位分
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 买家实际支付给卖家的金额，单位分
	ActualTotalFee int64 `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}
