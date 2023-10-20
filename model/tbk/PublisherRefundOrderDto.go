package tbk

// PublisherRefundOrderDto 结构体
type PublisherRefundOrderDto struct {
	// 维权编号，是当前订单发生维权退款的编号，非淘宝订单编号，如订单发生多次维权，则会产生多个维权编号
	RefundSuitId string `json:"refund_suit_id,omitempty" xml:"refund_suit_id,omitempty"`
	// 淘宝父订单编号(买家在淘宝后台显示的订单编号)
	TbTradeParentId string `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// 淘宝子订单编号
	TbTradeId string `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 订单创建时间
	TbTradeCreateTime string `json:"tb_trade_create_time,omitempty" xml:"tb_trade_create_time,omitempty"`
	// 订单结算时间
	EarningTime string `json:"earning_time,omitempty" xml:"earning_time,omitempty"`
	// 维权创建时间
	TkRefundTime string `json:"tk_refund_time,omitempty" xml:"tk_refund_time,omitempty"`
	// 维权完成时间
	TkRefundSuitTime string `json:"tk_refund_suit_time,omitempty" xml:"tk_refund_suit_time,omitempty"`
	// 订单更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 推广者角色(二方、三方)
	TkOrderRole string `json:"tk_order_role,omitempty" xml:"tk_order_role,omitempty"`
	// 结算金额(订单确认收货后的成交金额）
	TbTradeFinishPrice string `json:"tb_trade_finish_price,omitempty" xml:"tb_trade_finish_price,omitempty"`
	// 维权金额(买家申请维权退款的金额)
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 结算预估收入=结算金额*提成。以订单确认收货后的成交金额为基数，预估您可能获得的收入。
	PubShareFee string `json:"pub_share_fee,omitempty" xml:"pub_share_fee,omitempty"`
	// 应退还佣金(不含技术服务费和渠道专项服务费)
	TkCommissionFeeRefund string `json:"tk_commission_fee_refund,omitempty" xml:"tk_commission_fee_refund,omitempty"`
	// 应退还补贴(不含技术服务费和渠道专项服务费)
	TkSubsidyFeeRefund string `json:"tk_subsidy_fee_refund,omitempty" xml:"tk_subsidy_fee_refund,omitempty"`
	// 应退还佣金对应的技术服务费
	TkCommissionAlimmRefundFee string `json:"tk_commission_alimm_refund_fee,omitempty" xml:"tk_commission_alimm_refund_fee,omitempty"`
	// 应退还补贴对应的技术服务费
	TkSubsidyAlimmRefundFee string `json:"tk_subsidy_alimm_refund_fee,omitempty" xml:"tk_subsidy_alimm_refund_fee,omitempty"`
	// 应退还佣金对应的渠道专项服务费
	TkCommissionAgentRefundFee string `json:"tk_commission_agent_refund_fee,omitempty" xml:"tk_commission_agent_refund_fee,omitempty"`
	// 应退还补贴对应的渠道专项服务费
	TkSubsidyAgentRefundFee string `json:"tk_subsidy_agent_refund_fee,omitempty" xml:"tk_subsidy_agent_refund_fee,omitempty"`
	// 应退还预估收入：订单发生维权退款应退还的预估收入（佣金+补贴，含技术服务费和渠道专项服务费）
	ShowReturnFee string `json:"show_return_fee,omitempty" xml:"show_return_fee,omitempty"`
	// 维权创建(淘客结算回执) 4, 维权成功(淘客结算回执) 2, 维权失败(淘客结算回执) 3, 发生多次维权，待处理 11, 从淘客处补扣（钱已结给淘客） 等待扣款 12, 从淘客处补扣（钱已结给淘客） 扣款成功 13, 从卖家处补扣（钱已结给卖家） 等待扣款 14, 从卖家处补扣（钱已结给卖家） 扣款成功 15
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 渠道关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 会员关系id
	SpecialId int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
}
