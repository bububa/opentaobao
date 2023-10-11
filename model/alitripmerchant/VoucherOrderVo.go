package alitripmerchant

// VoucherOrderVo 结构体
type VoucherOrderVo struct {
	// 订单编号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// PMS订单号
	PmsCode string `json:"pms_code,omitempty" xml:"pms_code,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单状态描述
	OrderStatusDesc string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// 预订时间
	BookDate string `json:"book_date,omitempty" xml:"book_date,omitempty"`
	// 支付渠道
	PaymentChannel string `json:"payment_channel,omitempty" xml:"payment_channel,omitempty"`
	// 总价格
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 货币类型
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 扣款价格
	RefundCostAmount string `json:"refund_cost_amount,omitempty" xml:"refund_cost_amount,omitempty"`
	// 早餐类型
	BreakfastType string `json:"breakfast_type,omitempty" xml:"breakfast_type,omitempty"`
	// 早餐描述
	BreakfastDec string `json:"breakfast_dec,omitempty" xml:"breakfast_dec,omitempty"`
	// 订单人姓
	ContactLastName string `json:"contact_last_name,omitempty" xml:"contact_last_name,omitempty"`
	// 订单人名
	ContactFirstName string `json:"contact_first_name,omitempty" xml:"contact_first_name,omitempty"`
	// 订单人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 订单人邮箱
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 套餐卷名字
	VoucherName string `json:"voucher_name,omitempty" xml:"voucher_name,omitempty"`
	// 下单方式
	PlaceOrderType string `json:"place_order_type,omitempty" xml:"place_order_type,omitempty"`
	// 套餐卷Id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 区号
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 支付剩余多少秒
	PayRemainTime int64 `json:"pay_remain_time,omitempty" xml:"pay_remain_time,omitempty"`
	// 支付类型
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
}
