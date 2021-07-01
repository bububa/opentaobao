package westcrm

// UserConsumerVo 结构体
type UserConsumerVo struct {
	// 订单
	Orders []OrderVo `json:"orders,omitempty" xml:"orders>order_vo,omitempty"`
	// 消费总额
	TotalPay string `json:"total_pay,omitempty" xml:"total_pay,omitempty"`
	// 用户id
	IbUserId int64 `json:"ib_user_id,omitempty" xml:"ib_user_id,omitempty"`
}
