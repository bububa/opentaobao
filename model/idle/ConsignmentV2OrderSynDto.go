package idle

// ConsignmentV2orderSynDto 结构体
type ConsignmentV2orderSynDto struct {
	// 订单子状态
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 订单主状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 闲鱼订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 下单环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 不同的状态传递不同参数
	Attribute *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
}
