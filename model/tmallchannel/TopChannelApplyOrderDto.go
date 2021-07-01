package tmallchannel

// TopChannelApplyOrderDto 结构体
type TopChannelApplyOrderDto struct {
	// 申请单单号
	ChannelPurchaseApplyOrderNo string `json:"channel_purchase_apply_order_no,omitempty" xml:"channel_purchase_apply_order_no,omitempty"`
	// 渠道
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 分销商名称
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 审核状态：0-无审核，-审核不通过，2-审核中，3-审核通过，4-撤回
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 申请单状态：100-待提交，0-创建中，1-待审核，2-订单完结，3-订单关闭，4-审核已通过，5-审核已拒绝，6-审核拒绝待重新修改
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 交易类型
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 审核描述
	AuditDesc string `json:"audit_desc,omitempty" xml:"audit_desc,omitempty"`
	// 撤回的描述
	CancelDesc string `json:"cancel_desc,omitempty" xml:"cancel_desc,omitempty"`
	// 审核创建时间
	ApplyCreateTime string `json:"apply_create_time,omitempty" xml:"apply_create_time,omitempty"`
	// 申请单详情
	ApplyOrderDetail *TopChannelApplyOrderDetailDto `json:"apply_order_detail,omitempty" xml:"apply_order_detail,omitempty"`
	// 解析详情
	Schema string `json:"schema,omitempty" xml:"schema,omitempty"`
}
