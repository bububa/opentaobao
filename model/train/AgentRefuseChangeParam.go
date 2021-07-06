package train

// AgentRefuseChangeParam 结构体
type AgentRefuseChangeParam struct {
	// 扩展参数
	ExtendParam string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 改签申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 拒绝原因类型，0 其他、1 余票不足、2 已取票、3 已线下退票
	RefuseType int64 `json:"refuse_type,omitempty" xml:"refuse_type,omitempty"`
	// 代理商id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 订单id
	MainBizOrderId int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
}
