package alsc

// PageQueryAccountFlowsOpenInfo 结构体
type PageQueryAccountFlowsOpenInfo struct {
	// 储值账户id
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 储值账户流水id
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 储值相关的交易类型
	FlowType string `json:"flow_type,omitempty" xml:"flow_type,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 操作人ID
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 操作人
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 交易时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 外部交易单号id
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 理由
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 交易门店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 交易门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 创建者
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 更新者
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 外部支付订单id
	OuterPayId string `json:"outer_pay_id,omitempty" xml:"outer_pay_id,omitempty"`
	// 支付方式需要按照标准格式传入
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 交易后剩余总金额
	CurrentValue int64 `json:"current_value,omitempty" xml:"current_value,omitempty"`
	// 外部订单来源
	OrderSrc int64 `json:"order_src,omitempty" xml:"order_src,omitempty"`
	// 交易总金额，增加为正数，减少为负数
	OrderValue int64 `json:"order_value,omitempty" xml:"order_value,omitempty"`
	// 赠送金额
	GiftValue int64 `json:"gift_value,omitempty" xml:"gift_value,omitempty"`
	// 是否删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}
