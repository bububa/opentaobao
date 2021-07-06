package train

// AgentPayOrderParam 结构体
type AgentPayOrderParam struct {
	// 12306票号
	SequenceNo string `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
	// 12306支付url
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// 改签单号
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 主订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}
