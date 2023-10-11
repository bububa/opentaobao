package idle

// TenderPrePayCmd 结构体
type TenderPrePayCmd struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 预付款流水号
	PrePayNo string `json:"pre_pay_no,omitempty" xml:"pre_pay_no,omitempty"`
	// 预付款执行动作
	PrePayAction int64 `json:"pre_pay_action,omitempty" xml:"pre_pay_action,omitempty"`
}
