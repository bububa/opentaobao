package servicecenter

// TpFundsSendResultDo 结构体
type TpFundsSendResultDo struct {
	// 红包ID
	RedPacketsId string `json:"red_packets_id,omitempty" xml:"red_packets_id,omitempty"`
	// 红包发放时间，格式yyyy-MM-dd HH:mm:ss
	ReceiveTime string `json:"receive_time,omitempty" xml:"receive_time,omitempty"`
	// 金额，单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 资金权益类型。1：预付款红包；2：尾款红包
	FundsType int64 `json:"funds_type,omitempty" xml:"funds_type,omitempty"`
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 状态，true表示发放成功
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
